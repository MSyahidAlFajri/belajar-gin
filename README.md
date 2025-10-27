# belajar-gin

## Instalasi gin 
 ```bash
go get github.com/gin-gonic/gin
```
## Instalasi gorm

 ```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## Endpoint

### `GET /api/v1/users`

get semua data user

### `GET /api/users/:id`

get data user berdasarkan id

### `POST /api/v1/users`

tambah user 

### `PATCH /api/v1/users/:id`

update user

### `DELETE /api/v1/users/:id`

delete user berdasarkan id
