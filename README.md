# Go Steering
This library is intended to be a go implementation of OpenSteer. However this project doesn't work and is pretty much abandoned. 

##
Copy the wasm_exec.js file to your project's directory
`cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .`

Build the code and put it in main.wasm
`GOOS=js GOARCH=wasm go build -o main.wasm`

Install goexec if needed

`# install goexec: go get -u github.com/shurcooL/goexec`

Run a webserver from the command line

`$ goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'`
