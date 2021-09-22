package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Migrate(db *pgxpool.Pool) (bool, error) {
	eventQuery := `CREATE TABLE IF NOT EXISTS public.events
	(
		id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
		type character varying(255) COLLATE pg_catalog."default" NOT NULL,
		"time" character varying(255) COLLATE pg_catalog."default" NOT NULL,
		identity character varying(255) COLLATE pg_catalog."default" NOT NULL,
		CONSTRAINT events_pkey PRIMARY KEY (id)
	)`

	if _, err := db.Exec(context.Background(), eventQuery); err != nil {
		return false, err
	}

	return true, nil

}
