# go programming

## keywords 

package , func , import

## builtin functions

print, println 


## commands 

- to run go application

```
go run main.go
```

- to build

```
go build main.go
```

```
go build -o hello main.go
```

- list of all available os/arch

```
go tool dist list
```
GOOS/GOARCH

- cross compilation
```
 GOOS=linux GOARCH=s390x go build -o hello-release-ibm main.go
 ```

- stripe down the binary size for IBM bases machines 

```
 GOOS=linux GOARCH=s390x go build -ldflags="-s -w" -o hello-release-ibm main.go
 ```
