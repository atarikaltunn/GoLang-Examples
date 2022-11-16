# From Example 1: hello.go

In this example I have learned the creation and usage of a GoLang package.

## Some Important Commands:
-   go mod init <module_name>
    Creates a dependencies module that contains name of module, names of requirements, packages etc.
-   go run .
    Runs your code
-   go mod tidy
    Adds new module requirements and sums.


# From Example 2: greetings.go

- fundamental structure of funciton in go is explained. It should be like 
    ```
    func <function_name>(parameters) <return type>
    ```
- := operator is explained. 

    ```
    var message string
    message = fmt.Sprintf("Hi, %v. Welcome!", name)
    ```

    is equals to `message := fmt.Sprintf("Hi, %v. Welcome!", name)`.

    := operator helps us to declare and create a new variable.