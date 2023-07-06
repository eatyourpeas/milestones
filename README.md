# Milestones

## Childhood developmental milestones. A REST API in Golang.

There are lots of applications out there that seek to be a reference for children's developmental milestones. As a topic it is complicated because there are lots of different scoring systems, and these are all quite culturally dependent. For example, a measure in the Griffith's scales, developed in the 1950s for gross motor of an older child involved reported ability to jump onto a moving bus - completely appropriate for the time, but not applicable now. 

### Why Golang?

Golang largely because I have not used it before and intrigued by its versatility, speed and multithreading. In fact though, there is a lot of boiler plate I have learnt in setting up a simple REST API, for interfacing with the database, defining routes and schemas. And that is before authentication and so on. 

In any case, this is a functioning service in principle, and the concept of delivering developmental scales and subscales as an API is workable.

### Getting started

1. install [Golang](https://go.dev/)
2. Fork this repo
3. ```cd milestones```console
4. ```go build milestones```
5. ```./milestones```
