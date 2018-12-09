# tserver
Minimum Server written in golang

#### Structure of Directories:
- **Golang Server**:
    -  [webapp root] / src / github.com / tcarvi / tserver / main.go
- **Index file**:
    -  [webapp root] / webapp / [hostname] /index.html
- **API scripts**:
    -  [webapp root] / webapp / api / apiScripts.js
- **General CSS**:
    -  [webapp root] / webapp / layout / normalize.css
    -  [webapp root] / webapp / layout / boilerplate.css
- **Js scripts**:
    -  [webapp root] / webapp / [hostname] / js / scripts.js
- **CSS files**:
    -  [webapp root] / webapp / [hostname] / css /designFiles.css
- **ICO image**:
    -  [webapp root] / webapp / [hostname] / favicon.ico
- **SVG images**:
    -  [webapp root] / webapp / svg / images.svg
- **PNG images**:
    -  [webapp root] / webapp / png / images.png
- **JPG images**:
    -  [webapp root] / webapp / jpg / images.jpg
- **PDF files**:
    -  [webapp root] / webapp / pdf / files.pdf

#### Obs.:
- To run on localhost, you must edit your hosts file (/etc/hosts in linux):
```
127.0.0.1           localhost
::1                 localhost
[hostname].com      localhost:8080
```