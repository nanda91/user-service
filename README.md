# USER-SERVICE
============================

## Folder Structure Conventions

> Folder structure options and naming conventions for software projects

### A typical top-level directory layout

    .
    ├── config        # Layer yang akan bertugas meload atau init data dari config.json
    ├── delivery      # Layer yang akan bertugas sebagai presenter atau menjadi output dari aplikasi
    ├── lib           # Direktori core function
    ├── routes        # Direktori yang mendefinisikan route rest
    ├── model         # Layer yang mendefinisikan field
    ├── repository    # Layer yang menyimpan database handler. Querying, Inserting
    └── README.md

### Setup

- Make sure that you are using go 1.17
- The directory name should be `user-service` 
- setup `config.json`
- run `go mod tidy`

### Run App

- `go run main.go`
- using docker `docker-compose up`
 