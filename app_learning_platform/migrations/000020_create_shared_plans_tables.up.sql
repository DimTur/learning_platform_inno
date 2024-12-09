CREATE TABLE IF NOT EXISTS "shared_plans_users" (
  "id" SERIAL PRIMARY KEY,
  "plan_id" integer NOT NULL,
  "user_id" varchar(24) NOT NULL,
  "created_by" varchar(24) NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  CONSTRAINT fk_plan FOREIGN KEY ("plan_id") REFERENCES "plans" ("id") ON DELETE CASCADE,
  CONSTRAINT unique_plan_user UNIQUE ("plan_id", "user_id")
);
