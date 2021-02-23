package types

type Config struct {
	Title    string //标题
	From     string //发件人
	Subject  string //邮件标题
	Text     string //邮件正文
	HTML     string //邮件正文 html格式
	SMTP     string //smtp地址
	Port     int    //端口
	SMTPPort int    //发送端口
	Username string //邮件用户名
	Password string //邮件密码
	Font     string //字体
	Language string //语言
	I18n     map[string]struct {
		Download    string
		SendToEmail string
		Header      string
	}
	Scripts []string
}
