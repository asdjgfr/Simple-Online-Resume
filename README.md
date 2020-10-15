dev

use

```shell
sudo apt install texlive-xetex
```

```shell
bin/pandoc assets/resume.md --pdf-engine=xelatex -t latex -o assets/resume.pdf
```
查看支持字体
fc-list :lang=zh
填入CJKmainfont参数中
bin/pandoc assets/resume.md -o assets/resume.tmpl