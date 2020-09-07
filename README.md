# Photo API

Special feature here is storing photos and images. We need to provide our users with the possibility to search stored images based on attribute fields.

### Installation

Photo API requires Go 1.14+ to run and uses Go Modules as a vendoring tool.

Install the dependencies and start the server.

```sh
$ go mod vendor
$ go run main.go
```

### How to test

After running the app locally you have to wait until the indexer fetches all the pictures. 

```sh 
$ curl --location --request GET 'localhost:3000/search/author-twinconfusion'
```

You can replace the "author-index-twinfusion" following the next pattern:

- author-{author_name}
- camera-{camera_name}
- tag-{tag_name}

Always in camelcase and without blank spaces.

