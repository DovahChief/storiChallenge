CREATE TABLE IF NOT EXISTS public.report
(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255),
    balance numeric(4,2),
    avg_debit numeric(4,2),
    avg_credit numeric(4,2),
    creation_date timestamp with time zone NOT NULL DEFAULT now()
    );