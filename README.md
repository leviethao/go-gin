#Run
go install
go run server.go

#Test
POST http://localhost:8080/api/videos
body: {
    "title": "tile Cool",
    "description": "some description",
    "url": "https://google.com",
    "author": {
        "firstname": "haolv",
        "lastname": "le",
        "age": 10,
        "email": "leviethao2@gmail.com"
    }
}

header: Authorization: basic haolv:haolv123456 base64 encode

GET http://localhost:8080/api/videos

GET http://localhost:8080/view/videos