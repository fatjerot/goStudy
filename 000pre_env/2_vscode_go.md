# vscode + golang

1. 下载解压，然后就可以直接运行
   ```
   wget https://vscode.cdn.azure.cn/stable/b37e54c98e1a74ba89e03073e5a3761284e3ffb0/code-stable-1568209119.tar.gz
   tar -xvf code-stable-1568209119.tar.gz
   cd VSCode-linux-x64 && ./code
   ```
2. 安装插件：`Go`，`Chinese`，`vscode-Icons`，`filesize`，`Code Runner`，`GitLens`
3. 为了适配`go mod`，安装`gopls`
   ```
   go get golang.org/x/tools/gopls@latest
   ```
4. 配置golang环境：[ ctrl + , ] ，打开json格式的配置
   ```
   {
        "files.autoSave": "afterDelay",
        "editor.fontSize": 16,
        "workbench.iconTheme": "vscode-icons",
        "workbench.colorCustomizations" : {
            "terminal.background" : "#383737"
        },
        "search.followSymlinks": false,
        "terminal.integrated.lineHeight": 1.3,
        "terminal.integrated.letterSpacing": 0.1,
        "terminal.integrated.fontSize": 16,
        "terminal.integrated.fontFamily": "monospace",
        "go.inferGopath": false,
        "go.buildOnSave": "workspace",
        "go.lintOnSave": "package",
        "go.vetOnSave": "package",
        "go.buildTags": "",
        "go.buildFlags": [],
        "go.lintFlags": [],
        "go.vetFlags": [],
        "go.coverOnSave": false,
        "go.autocompleteUnimportedPackages": true,
        "go.useCodeSnippetsOnFunctionSuggest": true,
        "go.formatTool": "goreturns",
        "go.gopath": "/home/yjx/Repository/code/golang/package",
        "go.goroot": "/usr/local/go/go1.13",
        "go.docsTool": "gogetdoc",
        "go.gocodeAutoBuild": false,
        "go.gotoSymbol.includeImports": true,
        "go.useLanguageServer": true,
        "go.alternateTools": {
            "go-langserver": "gopls", 
        },
        "go.languageServerExperimentalFeatures": {
            "format": true,
            "autoComplete": true
        },
        "[go]": {
            "editor.snippetSuggestions": "none",
            "editor.formatOnSave": true,
            "editor.codeActionsOnSave": {
                "source.organizeImports": true
            }
        }
    }
   ```
5. vscode会自动提示装一些插件，安装即可。
6. 安装go的相关辅助命令（ctrl + shift + p --> go:install/update tools --> 全选组件安装）。==注意！golang版本升级后，需要update一下这些插件，最好把$GOPATH/bin下的插件都更新一下！==
7. 这个组件`golang.org/x/tools/cmd/gorename`安装失败需要手动处理
   ```
   cd /home/yjx/Repository/code/golang/package && mkdir -p src/golang.org/x/ && cd src/golang.org/x/
   git clone https://github.com/golang/tools.git tools
   go install golang.org/x/tools/cmd/gorename
   ```