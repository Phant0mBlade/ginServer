
## Gin Server

Boiler plate code for gin server.

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