




# WSLight

Unix command translator on Windows cmd (CLi)

----

<p align="center">
<strong><a href="#-idea">🔦 Idea</a></strong>
|
<strong><a href="#-installation">💺 Installation</a></strong>
|
<strong><a href="#-usage">🚀 Usage</a></strong>
|
<strong><a href="#-commands--features">📝 Commands</a></strong>
|
<strong><a href="#limitsimprovements">💭Limits/improvements</a></strong>
</p>

----

	
## 🔦 Idea
**Aim?**

Providing equivalent of bash (Unix cmd) on Window

> Powershell & cmd already exist!

Exactly, but if you are not very comfortable with it and you have more references with unix commands

> [WSL](https://itsfoss.com/install-bash-on-windows/) already exists  and provides more utilities!

Yes but it requires  admin privileges to enable it. 

If the command you want to launch is not available [see](#limitsimprovements)

**Use cases?**

 - CTF, pentest etc with remote shell on windows device
 - Learn some `cmd.exe` command (by enabling debug with `+x`)


 ## 💺 Installation
 
Clone the repo and download the dependencies locally:
```
git clone https://github.com/ariary/wslight.git
make before.build
```

To build wslight :
```
make build.wslight
```

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

see [available commands](#-commands--features) or Type `help` to get available commands

### Activate debug mode

Debug mode is useful to see which command is in fact launch on cmd (to see the translation). It is useful to debug your behaviour or just learn some command
```
> +x
```

Disable (by default) it with (`-x`)

## 📝 Commands & Features

 - `wslight` could understand pipe commands

List of available commands


| Unix command  | flag accepted|
|:--|:--|
| hostname ||
| pwd||
| rm |-r,-f|
| cp|-r |
|grep| -R, -i|
|ls|-l, -R, -a|
|tree||
|env||
|cd| -, ~|

*flags must be separated with spaces to be parsed correctly*

If the command you want to launch is not available [see](#limitsimprovements)

## 💭Limits/improvements

**⚠️ Only a set of Unix command is available. I will provide more as things progress. Do not hesitate to let me know which one you will be interested in**

*(Or a make a PR it is **simple**, just add a line in the `suggestions` slice (`cmd/wslight/main.go`) explaining what is the command, and  apply the command behavior in the `Translate` function (in `pkg/command/translate.go`)*

 - From now cli have no context (ie make a cd has no impact)

