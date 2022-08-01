CREATE TABLE "customers" (
	"id" SERIAL PRIMARY KEY,
	"email" VARCHAR(128) NOT NULL,
	"name" VARCHAR(64) NOT NULL,
	"firstname" VARCHAR(64) NOT NULL,
	"birthday" DATE NOT NULL,
	"driving_licence_number" VARCHAR(128) NOT NULL,
    UNIQUE ("email"),
    UNIQUE ("driving_licence_number")
)
;

COMMENT ON COLUMN "customers"."id" IS '';
COMMENT ON COLUMN "customers"."email" IS '320 max';
COMMENT ON COLUMN "customers"."name" IS '';
COMMENT ON COLUMN "customers"."firstname" IS '';
COMMENT ON COLUMN "customers"."birthday" IS '';
COMMENT ON COLUMN "customers"."driving_licence_number" IS '';

/* ALTER TABLE "customers" ADD UNIQUE ("email");
ALTER TABLE "customers" ADD UNIQUE ("driving_licence_number"); */
