# Gradient Avatar Service - GAS

Gradient Avatar Service is a service that generates avatars based on a user's name. GAS is inspired by Vercel's user avatar service. This project is fully written in Go and uses the Fiber framework for creating  REST API.

# Routes

    GET /avatar?name=<name>
        Returns an gradient avatar for the given name.
    
    GET /avatar?name=<name>&size=<size>
        Returns an gradient avatar for the given size and name.

**NOTE: by default, the size is 200x200.**

# Soon

- public API for generating avatars coming soon (for now, clone the repo and run `go run main.go`)


Author: [@n4ze3m](https://twitter.com/n4ze3m)