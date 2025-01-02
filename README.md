# go-movies-crud
CRUD API for movies built using go

# Download mux using the following command
`go get github.com/gorrila/mux`

# Working of CRUD API 

| Routes   | Endpoint     | Functions   | Method |
| -------- |:-------------| :-----------|:------ |
| GetAll   | /movies      | GetMovies   | GET    |
| GetById  | /movies/{id} | GetMovie    | GET    |
| Create   | /movies      | CreateMovie | POST   |
| Update   | /movies/{id} | UpdateMovie | PUT    |
| Delete   | /movies/{id} | DeleteMovie | DELETE |
