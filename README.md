
# WSLight

Unix command translator on Windows cmd (CLi)

##  Table of contents

 - [🔦 Idea](#-idea)
 - [💺 Installation](#-installation)
 - [🚀 Usage](#-usage)
	
## 🔦 Idea
**Aim?**

Providing equivalent of bash (Unix cmd) on Window

> Powershell & cmd already exist!

Exactly, but if you are not very comfortable with it and you have more references with unix commands

> [WSL](https://itsfoss.com/install-bash-on-windows/) already exists  and provides more utilities!

Yes but it requires  admin privileges to enable it. 

**⚠️ Only a set of Unix command is available. I will provide more as things progress. Do not hesitate to let me know which one you will be interested in**

*(Or a make a PR it is **simple**, just add a line in the `suggestions` slice (`cmd/wslight/main.go`) explaining what is the command, and  apply the command behavior in the `Translate` function (in `pkg/command/translate.go`)*


**Use cases?**

 - CTF, pentest etc with remote shell on windows device
 - Learn some `cmd.exe` command (by enabling debug with `+x`)


 ## 💺 Installation

 ## 🚀 Usage 
 
Move `wslight.exe` binary on the windows machine on which you have a shell

### Launch the cli
```
wslight.exe
```


### Launch equivalent of unix command in `cmd.exe`
Once the cli is launched
```
> <your_command> <your_argument>
```
As simple as that!

Type `help` to get available commands

### Activate debug mode

Debug mode is useful to see which command is in fact launch on cmd (to see the translation). It is useful to debug your behaviour or just learn some command
```
> +x
```

Disable (by default) it with (`-x`)

