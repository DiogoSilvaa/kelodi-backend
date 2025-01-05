# Remote Deployment

This folder contains scripts and configuration files for deploying the Kelodi API to a remote server.

## Folders

### [production](production)

- Contains configuration files and scripts for setting up and deploying the Kelodi API in a production environment.

## Files

### [api.service](production/api.service)

- **Systemd Service File**: Defines the systemd service for running the Kelodi API. Includes configuration for starting, stopping, and restarting the API service.

### [Caddyfile](production/Caddyfile)

- **Caddy Configuration**: Defines the Caddy web server configuration for reverse proxying requests to the Kelodi API and handling HTTPS.

### [01.sh](setup/01.sh)

- **Setup Script**: A shell script for setting up the production server. It installs necessary packages, sets up PostgreSQL, configures the firewall, and installs Caddy.
