<details>
    <summary>
        <strong>Only support *inux,windows user can use WSL to dev.</strong>
    </summary>
    仅支持*inux系统，windows可以使用wsl进行开发。
</details>


<h1 align="center">Simple Online Resume</h1>

<div align="center">
	A markdown to online and pdf program. Support download and send email.
    <br/>
    Markdown转在线和pdf的简历网站，支持下载和邮件发送。
</div>


![截图1](https://github.com/asdjgfr/resume/blob/master/screenshot/p1.png?raw=true)

# config.json

- `title` - Website title, also resume `pdf` file name. 网站标题同时也是`pdf`的文件名称，例如`xxx的个人简历`。
- `subject` - Email title. 邮件的标题。
- `text` - Email content text. 邮件的正文。
- `html` - Email content `html` format. 邮件的正文，`html`格式。
- `SMTP` - Your `smtp` server. e.g. smtp.example.com. smtp地址。
- `port` - Program port. 程序运行的端口。
- `SMTPPort` - `smtp` port. `smtp`的端口。
- `username` - Usually email address. 用户名，通常是邮箱地址。
- `password` - Password. 密码。
- `font` - font. 字体。
- `i18n` - `i18n`
- `language` - Language. 语言，和`i18n`中的对应。

You can use `fc-list` to search which font has been used on your `os`.

>  e.g. `fc-list :lang=zh`

# DEV

1. Create your own `config.json`.
2. Run `go mod vendor`
3. Use `texlive-xetex` to convert `md ` to `pdf`. So you should install`texlive-xetex`first.

   ```shell
   sudo apt install texlive-xetex
   ```

4. Then you can use `./run.sh` width dev mod.

   ```shell
   ./run.sh
   ```

# Build

1. Create your own `config.json`.

2. Run `go mod vendor`

3. And then:

   ```shell
   ./run.sh build
   ```

> Program will build in `build` folder.

# Use

1. Download lastest [release](https://github.com/asdjgfr/resume/releases)

2. install `texlive-xetex` on your server.

   ```shell
   sudo apt install texlive-xetex
   ```

3. Build `config.json` in program root folder and edit it.

4. Run:

   ```shell
   ./main
   ```

>  Change `assets/resume.md` and refresh browser. It can auto build resume without restart.