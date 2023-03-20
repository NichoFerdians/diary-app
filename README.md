# Diary-app

RESTful API menggunakan gin, gorm, jwt dan docker.

## Setup local development

### Tech Stack

- [Golang](https://golang.org/)
- [Gin](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [MySQL](https://www.mysql.com/)
- [JWT](https://github.com/dgrijalva/jwt-go)
- [Docker](https://www.docker.com/products/docker-desktop)


### How to run

- Run server:

    ```bash
    docker-compose up
    ```
    Setelah selesai API bisa diakses di localhost:8080, untuk detail endpoint bisa dilihat di resource postman.

### Resource

  [![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/9005645-afd717d5-58a2-49ea-a7de-12d2f7507f8d?action=collection%2Ffork&collection-url=entityId%3D9005645-afd717d5-58a2-49ea-a7de-12d2f7507f8d%26entityType%3Dcollection%26workspaceId%3D52de5606-528f-43b9-84fb-f9497ac41357)
- [Documentations](https://documenter.getpostman.com/view/9005645/2s93K1oz6X)
- DDL File => di file ./ddl.sql