CREATE TABLE "samples" (
    "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "title" TEXT NOT NULL,
    "content" TEXT NOT NULL,
    "photo" VARCHAR NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL
);
