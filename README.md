# ResourceServer
This module provides a handler to serve resource (files).
I started the project for fun and because of the problems between *Gorilla-Mux* and *FileServer*.
Use this command to get the module `go get github.com/StevenCyb/ResourceServer`.

## Example using http
```go
func main() {
  // ResourceServer.New(rootPath, cacheResources)
  // rootPath: root path to the resources
  // cacheResources: Define if resources should every time
  //                 be reade from file or hold in RAM
  http.Handle("/", ResourceServer.New("./static", true))

  // Basic http server stuff 
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}
```

## Example using Gorilla
```go
func main() {
	router := mux.NewRouter()

  // ResourceServer.New(rootPath, cacheResources)
  // rootPath: root path to the resources
  // cacheResources: Define if resources should every time
  //                 be reade from file or hold in RAM

  // If you know how deep your folder structure goes
  // e.g. if we know this can go up to two sub directories
  // "/folder1/folder2/index.html"
  router.Handle("/", ResourceServer.New("./static", true))
  router.Handle("/{l1}", ResourceServer.New("./static", true))
  router.Handle("/{l1}/{l2}", ResourceServer.New("./static", true))
  // Or use the "default" handler
  router.NotFoundHandler = ResourceServer.New("./static", true)

  // Basic http server stuff
  if err := http.ListenAndServe(":8080", router); err != nil {
    log.Fatal(err)
  }
}

```