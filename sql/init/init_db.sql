CREATE TABLE "cars" (
	"id" SERIAL PRIMARY KEY,
	"matriculation_number" VARCHAR(32) NOT NULL,
	"brand" VARCHAR(64) NOT NULL,
	"model" VARCHAR(64) NOT NULL,
	"year" INTEGER NOT NULL,
   	UNIQUE ("matriculation_number")
);

INSERT INTO "public"."cars" ("matriculation_number", "brand", "model", "year") VALUES ('AB123CD', 'Peugeot', '405', '1999');
INSERT INTO "public"."cars" ("matriculation_number", "brand", "model", "year") VALUES ('EF456GH', 'Renault', 'Megane', '1996');
INSERT INTO "public"."cars" ("matriculation_number", "brand", "model", "year") VALUES ('IJ789KL', 'Toyota', 'Yaris', '2001');
INSERT INTO "public"."cars" ("matriculation_number", "brand", "model", "year") VALUES ('MN147OP', 'Volkswagen', 'Polo', '1997');
INSERT INTO "public"."cars" ("matriculation_number", "brand", "model", "year") VALUES ('HI987JK', 'Fiat', '500', '2008');
INSERT INTO "public"."cars" ("matriculation_number", "brand", "model", "year") VALUES ('QR258ST', 'Peugeot', '406', '1998');
INSERT INTO "public"."cars" ("matriculation_number", "brand", "model", "year") VALUES ('UV369WX', 'Renault', 'Megane', '1997');
INSERT INTO "public"."cars" ("matriculation_number", "brand", "model", "year") VALUES ('ZA321BC', 'Toyota', 'Yaris', '2002');
INSERT INTO "public"."cars" ("matriculation_number", "brand", "model", "year") VALUES ('DE654FG', 'Volkswagen', 'Polo', '1998');
INSERT INTO "public"."cars" ("matriculation_number", "brand", "model", "year") VALUES ('LM753OP', 'Fiat', '500', '2008');



CREATE TABLE "customers" (
	"id" SERIAL PRIMARY KEY,
	"email" VARCHAR(128) NOT NULL,
	"name" VARCHAR(64) NOT NULL,
	"firstname" VARCHAR(64) NOT NULL,
	"birthday" DATE NOT NULL,
	"driving_licence_number" VARCHAR(128) NOT NULL,
    UNIQUE ("email"),
    UNIQUE ("driving_licence_number")
);

COMMENT ON COLUMN "customers"."id" IS '';
COMMENT ON COLUMN "customers"."email" IS '320 max';
COMMENT ON COLUMN "customers"."name" IS '';
COMMENT ON COLUMN "customers"."firstname" IS '';
COMMENT ON COLUMN "customers"."birthday" IS '';
COMMENT ON COLUMN "customers"."driving_licence_number" IS '';

INSERT INTO "public"."customers" ("email", "name", "firstname", "birthday", "driving_licence_number") VALUES ('jean.dupont@domain.com', 'Dupont', 'Jean', '1990-12-13', 'JEAND657054SM9IJ');
INSERT INTO "public"."customers" ("email", "name", "firstname", "birthday", "driving_licence_number") VALUES ('olivier.duchene@a-pro.fr', 'Duchene', 'Olivier', '1985-06-20', 'OLIVD657055SM9IJ');



CREATE TABLE "customer2car" (
	"id" SERIAL PRIMARY KEY,
	"customer_id" INTEGER NOT NULL,
	"car_id" INTEGER NOT NULL,
	UNIQUE ("car_id"),
	CONSTRAINT "customer_id" FOREIGN KEY ("customer_id") REFERENCES "customers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT "car_id" FOREIGN KEY ("car_id") REFERENCES "cars" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);

COMMENT ON COLUMN "customer2car"."id" IS '';
COMMENT ON COLUMN "customer2car"."customer_id" IS 'probably a customer for many car';
COMMENT ON COLUMN "customer2car"."car_id" IS 'a car for one customer';

INSERT INTO "public"."customer2car" ("customer_id", "car_id") VALUES ('2', '5');
INSERT INTO "public"."customer2car" ("customer_id", "car_id") VALUES ('2', '6');
INSERT INTO "public"."customer2car" ("customer_id", "car_id") VALUES ('1', '8');
INSERT INTO "public"."customer2car" ("customer_id", "car_id") VALUES ('1', '9');
INSERT INTO "public"."customer2car" ("customer_id", "car_id") VALUES ('1', '10');

