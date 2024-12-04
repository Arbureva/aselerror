# aselerror

this is a package designed for handling errors related to 'Exceptional API Response.'

in traditional response handling, it's common to determine whether an error contains
sensitive information before deciding to return it to the front end. Frequent use of
`if` and `switch` statements might not be the best approach. `aselerror` simplifies
this process by comparing error levels, significantly reducing the response time of
interfaces.



## Important Notes

- As `aselerror` is relatively simple in design, it is recommended to use it in conjunction with the `errors` package for stack trace information.
- `aselerror` is intended for exceptional API response situations. Under normal circumstances, it is advised to avoid using it as error messages should be concise and clear.



## Installation

```
go get github.com/Locter9001/aselerror
```


## Usage Example

you can check ths `error_test.go` file => [error_test.go](error_test.go)