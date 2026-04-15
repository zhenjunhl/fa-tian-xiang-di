package f

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/smtp"
	"os"
	"strings"
)

// EmailConfig 邮件服务器配置（必备字段）
type EmailConfig struct {
	SMTPHost    string // SMTP服务器地址
	SMTPPort    int    // SMTP服务器端口
	Username    string // 用户名
	Password    string // 密码
	FromAddress string // 发件人地址
	FromName    string // 发件人名称
	TLS         bool   // 是否使用TLS
}

// Email 邮件内容结构（必备字段）
type Email struct {
	To          []string // 收件人列表
	CC          []string // 抄送列表
	BCC         []string // 密送列表
	Subject     string   // 邮件主题
	Body        string   // 邮件正文
	IsHTML      bool     // 是否为HTML格式
	Attachments []string // 附件路径列表
}

// SendEmail 发送邮件方法
func (config *EmailConfig) SendEmail(email *Email) error {
	// 创建SMTP客户端
	auth := smtp.PlainAuth("", config.Username, config.Password, config.SMTPHost)
	
	// 构建邮件消息
	msg := &bytes.Buffer{}
	writer := multipart.NewWriter(msg)
	
	// 写入邮件头
	fmt.Fprintf(msg, "From: %s <%s>\r\n", config.FromName, config.FromAddress)
	fmt.Fprintf(msg, "To: %s\r\n", strings.Join(email.To, ", "))
	
	if len(email.CC) > 0 {
		fmt.Fprintf(msg, "Cc: %s\r\n", strings.Join(email.CC, ", "))
	}
	
	if len(email.BCC) > 0 {
		fmt.Fprintf(msg, "Bcc: %s\r\n", strings.Join(email.BCC, ", "))
	}
	
	fmt.Fprintf(msg, "Subject: %s\r\n", email.Subject)
	
	if email.IsHTML {
		fmt.Fprintf(msg, "Content-Type: multipart/mixed; boundary=%s\r\n\r\n", writer.Boundary())
		
		// 写入HTML部分
		part, err := writer.CreatePart(map[string][]string{
			"Content-Type": {"text/html; charset=utf-8"},
		})
		if err != nil {
			return err
		}
		part.Write([]byte(email.Body))
	} else {
		fmt.Fprintf(msg, "Content-Type: text/plain; charset=utf-8\r\n\r\n")
		msg.Write([]byte(email.Body))
		return smtp.SendMail(fmt.Sprintf("%s:%d", config.SMTPHost, config.SMTPPort), auth, config.FromAddress, getAllRecipients(email), msg.Bytes())
	}
	
	// 添加附件
	for _, attachment := range email.Attachments {
		err := addAttachment(writer, attachment)
		if err != nil {
			return err
		}
	}
	
	writer.Close()
	
	// 发送邮件
	return smtp.SendMail(fmt.Sprintf("%s:%d", config.SMTPHost, config.SMTPPort), auth, config.FromAddress, getAllRecipients(email), msg.Bytes())
}

// getAllRecipients 获取所有收件人（包括To、CC、BCC）
func getAllRecipients(email *Email) []string {
	recipients := make([]string, 0)
	recipients = append(recipients, email.To...)
	recipients = append(recipients, email.CC...)
	recipients = append(recipients, email.BCC...)
	return recipients
}

// addAttachment 添加附件
func addAttachment(writer *multipart.Writer, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	
	part, err := writer.CreatePart(map[string][]string{
		"Content-Type": {"application/octet-stream"},
		"Content-Disposition": {fmt.Sprintf("attachment; filename=%s", fileInfo.Name())},
	})
	if err != nil {
		return err
	}
	
	_, err = io.Copy(part, file)
	return err
}
