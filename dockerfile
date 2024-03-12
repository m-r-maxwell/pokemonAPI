FROM postgres:latest

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=pokemon

RUN mkdir -p /docker-entrypoint-initdb.d
COPY db/items/000items.sql /docker-entrypoint-initdb.d/000items.sql
COPY db/items/001items.sql /docker-entrypoint-initdb.d/001items.sql
COPY db/pokemon/000pokemon.sql /docker-entrypoint-initdb.d/000pokemon.sql
COPY db/pokemon/001pokemon.sql /docker-entrypoint-initdb.d/001pokemon.sql
COPY db/pokemon/002pokemon.sql /docker-entrypoint-initdb.d/002pokemon.sql
COPY db/moves/000moves.sql /docker-entrypoint-initdb.d/004moves.sql
COPY db/moves/001moves.sql /docker-entrypoint-initdb.d/005moves.sql

EXPOSE 5432
CMD ["postgres"]
