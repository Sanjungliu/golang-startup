CREATE TABLE "users" (
    "id" serial primary key,
    "name" varchar,
    "occupation" varchar,
    "email" varchar,
    "password_hashed" varchar,
    "avatar_file_name" varchar,
    "role" varchar,
    "created_at" timestamptz default (now()),
    "updated_at" timestamptz default (now())
);

CREATE TABLE "campaigns" (
    "id" serial primary key,
    "user_id" int,
    "name" varchar,
    "short_description" varchar,
    "description" text,
    "goal_amount" int,
    "current_amount" int,
    "perks" text,
    "backer_count" int,
    "slug" varchar,
    "created_at" timestamptz default (now()),
    "updated_at" timestamptz default (now())
);

CREATE TABLE "campaign_images" (
    "id" serial primary key,
    "campaign_id" int,
    "file_name" varchar,
    "is_primary" boolean,
    "created_at" timestamptz default (now()),
    "updated_at" timestamptz default (now())
);

CREATE TABLE "transactions" (
    "id" serial primary key,
    "campaign_id" int,
    "user_id" int,
    "amount" int,
    "status" varchar,
    "code" varchar,
    "payment_url" varchar,
    "created_at" timestamptz default (now()),
    "updated_at" timestamptz default (now())
);

ALTER TABLE "campaigns" ADD FOREIGN KEY ("user_id") REFERENCES "users"("id");
ALTER TABLE "campaign_images" ADD FOREIGN KEY ("campaign_id") REFERENCES "campaigns"("id");
ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users"("id");
ALTER TABLE "transactions" ADD FOREIGN KEY ("campaign_id") REFERENCES "campaigns"("id");