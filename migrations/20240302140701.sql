-- Create "bills" table
CREATE TABLE "public"."bills" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "customer_name" text NOT NULL,
  "amount" numeric NOT NULL DEFAULT 0,
  "description" text NOT NULL,
  "paid" boolean NOT NULL DEFAULT false,
  PRIMARY KEY ("id")
);
-- Create index "idx_bills_deleted_at" to table: "bills"
CREATE INDEX "idx_bills_deleted_at" ON "public"."bills" ("deleted_at");
-- Create "rooms" table
CREATE TABLE "public"."rooms" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "room_number" text NOT NULL,
  "type" text NOT NULL,
  "description" text NOT NULL,
  "capacity" bigint NOT NULL,
  "price" numeric NOT NULL DEFAULT 0,
  "is_booked" boolean NOT NULL DEFAULT false,
  PRIMARY KEY ("id")
);
-- Create index "idx_rooms_deleted_at" to table: "rooms"
CREATE INDEX "idx_rooms_deleted_at" ON "public"."rooms" ("deleted_at");
-- Create "bookings" table
CREATE TABLE "public"."bookings" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "room_id" bigint NULL,
  "user_id" bigint NULL,
  "check_in_date" timestamptz NULL,
  "check_out_date" timestamptz NULL,
  "total_cost" numeric NOT NULL DEFAULT 0,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_rooms_booked_dates" FOREIGN KEY ("room_id") REFERENCES "public"."rooms" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_bookings_deleted_at" to table: "bookings"
CREATE INDEX "idx_bookings_deleted_at" ON "public"."bookings" ("deleted_at");
