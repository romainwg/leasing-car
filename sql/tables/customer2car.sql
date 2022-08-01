CREATE TABLE "customer2car" (
	"id" SERIAL PRIMARY KEY,
	"customer_id" INTEGER NOT NULL,
	"car_id" INTEGER NOT NULL,
	UNIQUE ("car_id"),
	CONSTRAINT "customer_id" FOREIGN KEY ("customer_id") REFERENCES "customers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT "car_id" FOREIGN KEY ("car_id") REFERENCES "cars" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
)
;
COMMENT ON COLUMN "customer2car"."id" IS '';
COMMENT ON COLUMN "customer2car"."customer_id" IS 'probably a customer for many car';
COMMENT ON COLUMN "customer2car"."car_id" IS 'a car for one customer';
