CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "isVerified" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "shops" (
  "id" SERIAL PRIMARY KEY,
  "shop_name" varchar NOT NULL,
  "shop_type" varchar NOT NULL,
  "address" varchar NOT NULL,
  "user_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "inventories" (
  "id" SERIAL PRIMARY KEY,
  "item_name" varchar NOT NULL,
  "quantity" int NOT NULL,
  "cost_price" int NOT NULL,
  "selling_price_standard" int NOT NULL,
  "status" varchar NOT NULL DEFAULT 'available',
  "shop_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sales" (
  "id" SERIAL PRIMARY KEY,
  "item_id" int NOT NULL,
  "quantity" int NOT NULL,
  "selling_price_actual" int NOT NULL,
  "payment_type" varchar NOT NULL,
  "shop_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "expenses" (
  "id" SERIAL PRIMARY KEY,
  "particular" varchar NOT NULL,
  "amount" int NOT NULL,
  "recipient" varchar NOT NULL,
  "payment_type" varchar NOT NULL,
  "shop_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "payouts" (
  "id" SERIAL PRIMARY KEY,
  "particular" varchar NOT NULL,
  "amount" int NOT NULL,
  "recipient" varchar NOT NULL,
  "payment_type" varchar NOT NULL,
  "shop_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "banks" (
  "id" SERIAL PRIMARY KEY,
  "bank_name" varchar NOT NULL,
  "icon" varchar NOT NULL,
  "account_number" varchar NOT NULL,
  "shop_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tokens" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "access_token" varchar,
  "refresh_token" varchar
);

CREATE INDEX ON "users" ("email");
CREATE INDEX ON "shops" ("shop_name");
CREATE INDEX ON "tokens" ("user_id");
CREATE INDEX ON "banks" ("account_number");

ALTER TABLE "tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "shops" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "inventories" ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("id");
ALTER TABLE "sales" ADD FOREIGN KEY ("item_id") REFERENCES "inventories" ("id");
ALTER TABLE "sales" ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("id");
ALTER TABLE "expenses" ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("id");
ALTER TABLE "payouts" ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("id");
ALTER TABLE "banks" ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("id");

ALTER TABLE "shops" ADD CONSTRAINT "shop_name_shop_type_address_user_id_key" UNIQUE ("shop_name", "shop_type", "address", "user_id");
ALTER TABLE "banks" ADD CONSTRAINT "bank_name_account_number_shop_id_key" UNIQUE ("bank_name", "account_number", "shop_id");
ALTER TABLE "inventories" ADD CONSTRAINT "item_name_shop_id_key" UNIQUE ("item_name", "shop_id");
