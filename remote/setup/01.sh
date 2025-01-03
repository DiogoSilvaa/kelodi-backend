set -eu

TIMEZONE="Europe/London"

USERNAME="kelodi"

read -p "Enter password for kelodi DB user: " DB_PASSWORD

export LC_ALL=en_US.UTF-8

add-apt-repository --yes universe

# Update software packages
apt update
apt --yes -o Dpkg::Options::="--force-confnew" upgrade

# Create a new user and force user to set a password first time they log in
useradd --create-home --shell "/bin/bash" --groups sudo "${USERNAME}"
passwd --delete "${USERNAME}"
chage --lastday 0 "${USERNAME}"

# Copy SSH keys from root user to new user
rsync --archive --chown=${USERNAME}:${USERNAME}/root/.ssh /home/${USERNAME}

# Set up firewall to allow SSH, HTTP and HTTPS
ufw allow 22 
ufw allow 80/tcp
ufw allow 443/tcp
ufw --force enable

# Install up fail2ban
apt --yes install fail2ban

# Install migrate CLI tool
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
mv migrate.linux-amd64 /usr/local/bin/migrate

# Install PostgreSQL
apt --yes install postgresql

# Set up PostgreSQL
sudo -i -u postgres psql -d kelodi -c "CREATE ROLE kelodi WITH LOGIN PASSWORD '${DB_PASSWORD}'"
sudo -i -u postgres psql -c "CREATE DATABASE kelodi OWNER kelodi"
sudo -i -u postgres psql -d kelodi -c "CREATE EXTENSION IF NOT EXISTS citext"
echo "KELODI_DB_DSN='postgres://kelodi:${DB_PASSWORD}@localhost/kelodi'" >> /etc/environment

# Install Caddy
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https curl
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy

echo "Script complete! Rebooting..."
reboot