# Migrations

This folder contains the database migration scripts for the Kelodi API. Each migration is responsible for creating, modifying, or deleting database schema objects.

## Naming Convention

Migration files follow a sequential naming convention with a brief description of the feature or change they implement. For example:

- `000001_feat_properties.up.sql`
- `000001_feat_properties.down.sql`

## Not Null Columns

All columns are defined as `NOT NULL` to ensure data integrity and simplify Golang interaction.
