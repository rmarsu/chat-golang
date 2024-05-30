Chat using Golang

Get Rooms

 curl "http://localhost:8080/ws/joinRooms"

Create Room

  curl -d '{"id": "1" ,"name":"smth"}' -H "Content-Type: application/json" -X POST http://localhost:8080/ws/createRoom

Join Room

  wscat -c "ws://localhost:8080/ws/joinRoom/1?clientId=1&username=name1"

Get Users in Room

  curl "http://localhost:8080/ws/getClients/1"

Sign Up

  curl -d '{"username":"user" ,"email":"user@email.com","password":"smth"}' -H "Content-Type: application/json" -X POST http://localhost:8080/auth/sign-up

Sign In

  curl -d '{"email":"user@email.com","password":"smth"}' -H "Content-Type: application/json" -X POST http://localhost:8080/auth/sign-in
