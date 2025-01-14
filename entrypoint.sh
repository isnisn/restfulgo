#!/bin/sh

# Check if PostgreSQL is ready
until pg_isready -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" -U "$POSTGRES_USER"; do
  echo "Waiting for Postgres to be ready..."
  sleep 5
done

echo "Postgres is ready. Starting the application..."

# Execute the main Go application
exec "$@"
