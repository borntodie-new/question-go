CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "username" varchar UNIQUE NOT NULL,
                         "hashed_password" varchar NOT NULL,
                         "nickname" varchar,
                         "avatar" varchar DEFAULT './static/images/avatar/default.jpg',
                         "status" bool DEFAULT true,
                         "is_super" bool DEFAULT false,
                         "created_at" timestamptz NOT NULL DEFAULT (now()),
                         "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "profiles" (
                            "id" bigserial PRIMARY KEY,
                            "user_id" bigint,
                            "real_name" varchar,
                            "gender" int DEFAULT 0,
                            "quote" varchar,
                            "address" varchar,
                            "created_at" timestamptz NOT NULL DEFAULT (now()),
                            "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

COMMENT ON COLUMN "profiles"."gender" IS '1:男, 2:女, 0:未知';

ALTER TABLE "profiles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
