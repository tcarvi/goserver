# tserver
Minimum Server written in golang

#### Structure of Directories:
- **Golang Server**:
    - [server_root] / go / src / github.com / [Your_Github_User] / tserver / main.go
- **Webapp files**:
    - [server_root] / webapp / github.com / [Your_Github_User] / [ web_app_folders ]
    - For [web_app_folders] pattern, follow instructions from https://github.com/tcarvi/localhost

#### Container and/or Virtual machine configurations:
- You must have **golang tools** installed
- **$GOPATH** must be set to [server_root]/go
- **$GOBIN** must be set to $GOPATH/bin
- **$GOBIN** must be in your PATH variable
- Example of commands:
```
# Golang workspace variables  
export GOPATH=[server_root]/go
export GOBIN=$GOPATH/bin

# Path for golang development tools
# Verify if /usr/go/bin is set. If not, add it.
export PATH=/usr/local/go/bin:$PATH
# Make your golang projects to be executed outside its src directory.
export PATH=$GOBIN:$PATH
```  

#### Server Initialization:
- Go to tserver.go directory
- Execute: ```go install```
- Inside the root of the webapp directory, execute: ```$ tserver```

#### Test on localhost:
- Add a line for your [hostname], in your hosts file (/etc/hosts in linux):
```
127.0.0.1           localhost
::1                 localhost
[hostname].com      localhost:8080
```
- Go to [server_root]/webapp/github.com/[Your_Github_User]/ 
- Clone a WebApp template inside this directory: git clone https://github.com/tcarvi/localhost.git
- Follow instructions from https://github.com/tcarvi/localhost