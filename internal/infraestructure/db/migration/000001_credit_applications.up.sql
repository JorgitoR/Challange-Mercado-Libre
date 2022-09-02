CREATE TABLE "user_credits" (
  "user_id" bigserial PRIMARY KEY,
  "amount_total" bigint NOT NULL,
  "cant" bigint
);

CREATE TABLE "user_loans" (
  "id" bigserial PRIMARY KEY,
  "date" timestamp,
  "ammount" bigint NOT NULL,
  "user_id" bigint,
  "installment" numeric,
  "target" text
);

CREATE TABLE "debt_payments" (
  "id" bigserial PRIMARY KEY,
  "loan_id" bigint,
  "debt" bigint
);

CREATE TABLE "credit_applications" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "amount" bigint,
  "term" numeric,
  "rate" numeric,
  "target" varchar,
  "date" timestamp
);

CREATE TABLE "target" (
  "id" bigint PRIMARY KEY,
  "target" text
);

CREATE INDEX ON "user_credits" ("user_id");

CREATE INDEX ON "user_loans" ("user_id");

CREATE INDEX ON "debt_payments" ("loan_id");

CREATE INDEX ON "credit_applications" ("user_id");

COMMENT ON COLUMN "user_loans"."ammount" IS 'must be positive';

ALTER TABLE "user_loans" ADD FOREIGN KEY ("user_id") REFERENCES "credit_applications" ("user_id");

ALTER TABLE "debt_payments" ADD FOREIGN KEY ("loan_id") REFERENCES "user_loans" ("id");

ALTER TABLE "credit_applications" ADD FOREIGN KEY ("user_id") REFERENCES "user_credits" ("user_id");
