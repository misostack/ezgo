# Overview

## Requirements

1. A website about programming with GOlang
2. Must be built by GOlang without any built in framework
3. Acceptable user interface with bootstrap
4. Allow admin creates or updates content in the back office.
5. Support search, filter, pagination

... to be updated

## Details

### A. Design

1. Web servers
2. Routers
3. Static assets : images, css, javascripts
4. Template parsing
5. Development tools : sass, sass watcher, webpack with javascript, html templates
6. Data connecter
7. Forms
8. Authentication
9. CRUD
10. Mailer system
11. Secured with Recaptcha

#### Breaking down

**Architectural style : ROA ( resource oriented architecture)**

System will includes:

1. API Server : Golang
2. Web Admin Server : ReactJS
3. Web Server : Golang

##### Backend

> Routes: /api/* : all resource's URI will begin with "/api/"

##### Frontend

> Routes: /* : all the URI which does not begin with "/api/" is web pages

> Routes: /admin/* : all the URI which begin by admin is a ReactJS app for admin

### B. Deployment

1. Storing static assets on Amazon S3
2. Deploy go on heroku for development
3. CI & CD with heroku

#### Makefiles for go developers

- https://tutorialedge.net/golang/makefiles-for-go-developers/
- https://sohlich.github.io/post/go_makefile/
- https://github.com/heroku/go-getting-started

### C. Release

1. Setup VPS on Vultr
2. Setup CI & CD with Circle CI or Travis Github
3. Storing static assets on Amazon S3
4. Load balanced with 2 servers

### D. Setup

1. Config and getting start

```bash
Eg: your enviroment is "local"
export EZGO_ENV=local
cp .env .env.local
# Local
go run ezgo.go
# Others for running just use
EZGO_ENV=development && ./ezgo
```

# References

## Tools

1. Code editor : VS Code
2. Extensions:

- https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go
- https://marketplace.visualstudio.com/items?itemName=mikestead.dotenv

## Required Packages

1. godotenv

- Link : https://github.com/joho/godotenv
- Shortdesc : Package godotenv is a go port of the ruby dotenv library (https://github.com/bkeepers/dotenv)
- Usage:

```go
go get github.com/joho/godotenv
```
- Others : https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66

2. flag

- Link : https://golang.org/pkg/flag/
- Shortdesc : Package flag implements command-line flag parsing.
- Usage:

```go
import "flag"
var ip = flag.Int("flagname", 1234, "help message for flagname")
```

## Funny things

> There are boolean constants, rune constants, integer constants, floating-point constants, complex constants, and string constants. Rune, integer, floating-point, and complex constants are collectively called numeric constants.

- https://golang.org/ref/spec#Constants

> Go does not support struct constants