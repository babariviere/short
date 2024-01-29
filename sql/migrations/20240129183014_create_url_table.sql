-- Create "urls" table
CREATE TABLE "public"."urls" ("id" serial NOT NULL, "hash" text NOT NULL, "long_url" text NOT NULL, PRIMARY KEY ("id"));
-- Create index "urls_hash_key" to table: "urls"
CREATE UNIQUE INDEX "urls_hash_key" ON "public"."urls" ("hash");
