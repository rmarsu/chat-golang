CREATE TABLE "users" (
     "id" bigserial PRIMARY KEY,
     "username" varchar NOT NULL unique,
     "email" varchar NOT NULL unique,
     "password" varchar NOT NULL
)