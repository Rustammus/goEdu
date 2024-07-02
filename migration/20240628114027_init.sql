-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.users
(
    uuid   uuid DEFAULT GEN_RANDOM_UUID() NOT NULL PRIMARY KEY ,
    name   VARCHAR(100) NOT NULL,
    phone  VARCHAR(100),
    email  VARCHAR(100) NOT NULL UNIQUE ,
    password_hash VARCHAR(200) NOT NULL,
    crystals INT DEFAULT 0 NOT NULL,
    is_blocked BOOLEAN DEFAULT false NOT NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP(0),
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP(0)
);

CREATE TABLE public.tasks
(
    uuid   uuid DEFAULT GEN_RANDOM_UUID() NOT NULL PRIMARY KEY,
    user_uuid uuid NOT NULL,
    status VARCHAR(100),
    message VARCHAR(200),
    user_upload_link VARCHAR(100) NULL,
    user_download_link VARCHAR(100) NULL,
    server_upload_link VARCHAR(100) NULL,
    server_download_link VARCHAR(100) NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP(0),
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP(0),
    FOREIGN KEY (user_uuid) REFERENCES public.users(uuid) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.tasks;
DROP TABLE public.users;
-- +goose StatementEnd
