# Clean Architecture
This repository contains example code for [https://sachinsaini.dev/clean-architecture-in-practice/](https://sachinsaini.dev/clean-architecture-in-practice/)

## directory structure
 - `cmd/api` contains entry point for the app
 - `pkg/models/user` contains user model and relevant interfaces
 - `pkg/user/`
   - `delivery/http` contains http handlers
   - `repository/mysql` contains MySQL repository
   - `usecase/` contains usecase layer

Feel free to open PRs or issues to suggest changes :)