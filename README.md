
## Gin Server

Boiler plate code for gin server.
https://go.dev/doc/tutorial/web-service-gin

### Code organisation

How to set and unset env varibles, https://go.dev/doc/code

https://go.dev/doc/effective_go

Use tabs in this projects as that is what is recommended.

https://go.dev/doc/comment

Structuring is a bit tricky, you import everything from inside a folder not from inside a file. And a single folder can have a single package?

For example, I can structure my schemas in 2 ways,
- Say, we define 2 files, `album.go` and `user.go`. Both of them are inside the folder `schema`.
  - Then both of them have to necessarily have the same package name, they CANNOT be like `package albumscheam` and `package userschema`. A common package name is required.
  - The problem with this approach is that if I want to import say only `Album` I cannot do that. It will import everything inside the folder.
- So, the better way (atleast for now) is to compose them inside their respective folders.
  - we will have 2 schemas as, `schemas/album/album.go` and `schemas/user/user.go`
  - this also looks much cleaner.

And, I believe this is what we will do with every component. Say, we have an album and user controller, their paths will be something like, `controller/album/abc.go` and `controller/user/abc.go`

Though, this should not be necessary for tools like routers.

### Auth middleware

https://gin-gonic.com/docs/examples/using-basicauth-middleware/

### DI

https://golangforall.com/en/post/dependency-injection.html
https://github.com/uber-go/fx
https://github.com/uber-go/dig

### Logger abstraction

NOTE: Logs increase response time

1. Install zap and zerolog
2. create a file where you define these loggers, say as a variable
3. 


### Multi Modules
https://go.dev/doc/tutorial/workspaces

https://gowebexamples.com/
https://gowebexamples.com/basic-middleware/
https://gowebexamples.com/advanced-middleware/
https://gobyexample.com/closures


### Sync Once

`sync.Once` awesome

https://victoriametrics.com/blog/go-sync-once/ - Show race condition
https://www.slingacademy.com/article/the-power-of-sync-once-for-one-time-initialization-in-go/ - Says race condition not possible

### Zero Allocation

what is zero allocation and no reflection JSON
https://medium.com/@muroon/the-reason-why-zap-and-zerolog-make-it-zero-allocation-3ee8f69f660
https://medium.com/@ksandeeptech07/zero-allocation-in-go-ce29e6a9ffdc
