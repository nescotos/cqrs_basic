DROP TABLE IF EXISTS "tweets";

CREATE TABLE "tweets" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "body" text NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
  "deleted_at" TIMESTAMP,
  CONSTRAINT "pk_tweets" PRIMARY KEY ("id")
)