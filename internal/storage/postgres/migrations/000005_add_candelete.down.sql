begin;

alter table if exists "pastes" drop column canDelete;

commit;