# Alt Showcase Server

Alterantive server implementation for showcase app
Uses GoLang instead of NodeJS

## Get Started

### Locally

* Pull down git repo:
    git pull https://github.com/colin-hanbury/alt_showcase_server

* Start up docker engine
    * Open Docker Desktop on PC or Mac
    * https://docs.docker.com/desktop/


* Build and run server and db
    * Execute the following cmd in the root dir of the project you pulled down:
    docker compose up

* Verify locally with 200 response and welcome message:
    http://localhost:8080/welcome