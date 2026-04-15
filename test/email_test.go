package test

import (
	"fmt"
	"testing"

	"github.com/zhenjunhl/fa-tian-xiang-di/f"
)
func TestSendEmail1(t *testing.T) {
	fmt.Printf("测试发送邮件1\n")
}
func TestSendEmail(t *testing.T) {
	// 配置QQ邮箱服务器
	// 注意：使用QQ邮箱需要先在邮箱设置中开启SMTP服务并获取授权码
	// 获取授权码步骤：
	// 1. 登录QQ邮箱
	// 2. 进入"设置" -> "账户"
	// 3. 找到"POP3/IMAP/SMTP/Exchange/CardDAV/CalDAV服务"
	// 4. 开启"SMTP服务"
	// 5. 按照提示发送短信验证
	// 6. 获取生成的授权码
	fmt.Printf("QQ邮箱授权码\n")
	config := &f.EmailConfig{
		SMTPHost:    "smtp.qq.com",
		SMTPPort:    587,
		Username:    "503153081@qq.com", // QQ邮箱地址
		Password:    "", // QQ邮箱授权码（不是登录密码）
		FromAddress: "503153081@qq.com", // 发件人邮箱地址
		FromName:    "zhen jun", // 发件人名称
		TLS:         true, // QQ邮箱需要使用TLS
	}
	
	// 构建邮件
	email := &f.Email{
		To:          []string{"zhenjun.gm@gmail.com"},
		// CC:          []string{"cc@example.com"},
		// BCC:         []string{"bcc@example.com"},
		Subject:     "测试邮件",
		Body:        "<h1>Hello!</h1><p>这是一封测试邮件</p>",
		IsHTML:      true,
		Attachments: []string{},
	}
	
	// 注意：实际运行此测试需要填写正确的SMTP配置
	// 由于涉及真实邮件发送，这里注释掉实际发送代码
	// 仅测试结构体和参数的正确性
	fmt.Println("邮件配置和内容构建成功")
	fmt.Printf("SMTP服务器: %s:%d\n", config.SMTPHost, config.SMTPPort)
	fmt.Printf("发件人: %s <%s>\n", config.FromName, config.FromAddress)
	fmt.Printf("收件人: %s\n", email.To)
	fmt.Printf("抄送: %s\n", email.CC)
	fmt.Printf("密送: %s\n", email.BCC)
	fmt.Printf("主题: %s\n", email.Subject)
	fmt.Printf("是否HTML: %v\n", email.IsHTML)
	fmt.Printf("附件数量: %d\n", len(email.Attachments))
	
	// 如需测试实际发送，请取消下面的注释并填写正确的SMTP配置
	// /*
	err := config.SendEmail(email)
	if err != nil {
		t.Errorf("发送邮件失败: %v", err)
	} else {
		fmt.Println("邮件发送成功")
	}
	// */
}
