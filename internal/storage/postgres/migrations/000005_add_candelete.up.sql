begin;

alter table if exists "pastes" add column IF NOT EXISTS canDelete boolean not null;

commit;