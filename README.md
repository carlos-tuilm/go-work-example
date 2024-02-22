# Go modules tutorial

This is an example I made using the go modules to understand better how packages and dependencies are organized in a structured manner.

The important conclusion from this exercise is that this allows to more easily use dependencies from other packages, while also making the code more understandable and easier to implement newer functions to it. 

In this example, the workspace follows this structure:

- Workspace
    - mymodule
        - mypackage
            mypackage.go
        - main.go
        - go.mod
        - go.sum
    - example
        - hello
            - reverse
                - int.go
                - reverse.go
                - ...
            - ...
- go.work
- README.file

Where mymodule module is a package that I created. Inside of it I injected my own dependecies from cobra to demonstrate the use of the go get function in each module. 

From the git cloned repository titled example, we exported only an specific module located in example/hello/reverse (included one title int.go that we created). If we want to add another external module to our own module we would need to point it to the go.work file.

## References
Read them in order

## Digital Ocean tutorial
Source: https://www.digitalocean.com/community/tutorials/how-to-use-go-modules

## Go dev
Source: https://go.dev/doc/tutorial/workspaces#create_folder