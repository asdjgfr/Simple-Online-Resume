const downloadBtn = document.querySelector("#download");
const sendEmailBtn = document.querySelector("#send-to-email");
const toolBar = document.querySelector("#tool-bar");
const header = document.querySelector("#header");
const emailReg = /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/

downloadBtn.addEventListener("click", () => {
    const dlDom = document.createElement("a");
    dlDom.download = document.title + ".pdf";
    dlDom.href = "/assets/resume.pdf";
    dlDom.click();
});

sendEmailBtn.addEventListener("click", clickEmailBtn);

async function clickEmailBtn() {
    const p = prompt("请输入邮箱地址：");
    if (p === null) {
        return false;
    }

    if (emailReg.test(p)) {
        this.setAttribute("disabled", "disabled")
        this.innerText = "邮件发送中..."
        let res = {}
        try {
            res = await sendEmail(p);
        } catch (e) {
            alert(`错误：${e}`)
        }
        this.removeAttribute("disabled")
        if (res.status === 0 || res.status === 1) {
            alert(res.msg)
        }
        this.innerText = "重新发送"
    } else {
        alert("请输入正确的邮箱地址！名称允许汉字、字母、数字，域名只允许英文域名。");
        await clickEmailBtn()
    }
}

function sendEmail(address) {
    return fetch("/api/send-email", {
        body: `address=${address}`,
        cache: 'no-cache',
        headers: {
            'Content-type': 'application/x-www-form-urlencoded; charset=utf-8'
        },
        method: 'POST',
    }).then(res => res.json());
}

window.addEventListener("scroll", changeHeaderAndToolBar);

changeHeaderAndToolBar();

function changeHeaderAndToolBar(e) {
    const above100 = document.body.getBoundingClientRect().top < -100
    toolBar.style.bottom = above100 ? 0 : -60 + "px";
    if (e) {
        header.style.opacity = "0";
    }
}

