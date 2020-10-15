dev

use

```shell
sudo apt install texlive-xetex
```

```shell
bin/pandoc assets/resume.md --pdf-engine=xelatex -t latex -o assets/resume.pdf
```

bin/pandoc assets/resume.md -c pandoc.css -o assets/resume.html