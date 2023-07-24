CREATE TABLE "users" (
  "uid" bigserial PRIMARY KEY NOT NULL,
  "username" VARCHAR(255) NOT NULL,
  "password" VARCHAR(255) NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "address" VARCHAR(255) NOT NULL,
  "mobile_no" VARCHAR(10) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
  "proid" bigserial PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "description" VARCHAR(255) NOT NULL,
  "category" VARCHAR(255) NOT NULL,
  "price" DECIMAL(10,2) NOT NULL,
  "stock" INT NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "payments" (
  "pid" bigserial PRIMARY KEY,
  "order_id" bigint NOT NULL,
  "amount" DECIMAL(10,2) NOT NULL,
  "payment_type" VARCHAR(255) NOT NULL,
  "status" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "order_details" (
  "oid" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "total" DECIMAL(10,2) NOT NULL,
  "payment_id" bigint NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "order_items" (
  "oiid" bigserial PRIMARY KEY,
  "order_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "users" ("mobile_no");

CREATE INDEX ON "products" ("name");

CREATE INDEX ON "products" ("category");

CREATE INDEX ON "products" ("stock");

CREATE INDEX ON "products" ("name", "category");

CREATE INDEX ON "products" ("name", "stock");

CREATE INDEX ON "products" ("name", "category", "stock");

CREATE INDEX ON "payments" ("order_id");

CREATE INDEX ON "order_details" ("user_id");

CREATE INDEX ON "order_items" ("order_id");

CREATE INDEX ON "order_items" ("product_id");

CREATE INDEX ON "order_items" ("order_id", "product_id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("uid");

ALTER TABLE "payments" ADD FOREIGN KEY ("order_id") REFERENCES "order_details" ("oid");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "order_details" ("oid");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("proid");