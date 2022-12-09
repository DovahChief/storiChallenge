CREATE TABLE IF NOT EXISTS public.transaction
(
    id SERIAL PRIMARY KEY,
    transaction_id bigint NOT NULL,
    transaction_date timestamp without time zone NOT NULL,
    transaction_amount numeric(13,2),
    creation_date timestamp without time zone NOT NULL DEFAULT now(),
    report_id integer NOT NULL,
    CONSTRAINT fk_report FOREIGN KEY (report_id)
    REFERENCES public.report (id) MATCH SIMPLE
                            ON UPDATE NO ACTION
                            ON DELETE NO ACTION
    );