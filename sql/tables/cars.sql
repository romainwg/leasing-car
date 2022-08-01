CREATE TABLE "cars" (
	"id" SERIAL PRIMARY KEY,
	"matriculation_number" VARCHAR(32) NOT NULL,
	"brand" VARCHAR(64) NOT NULL,
	"model" VARCHAR(64) NOT NULL,
	"year" INTEGER NOT NULL,
   	UNIQUE ("matriculation_number")
)
;