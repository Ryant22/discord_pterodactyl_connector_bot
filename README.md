# Discord Pterodactyl Connector

[![Go](https://img.shields.io/badge/Go-1.18%2B-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A Discord bot written in Go for managing and monitoring Pterodactyl servers directly from Discord.

---

## 🚀 Features
- Start, stop, and restart Pterodactyl servers
- View server status and resource usage
- Secure authentication
- Easy setup and deployment

## 📦 Installation

### Clone the Repository or Download
```powershell
git clone https://github.com/yourusername/discord_pterodactyl_connector.git
cd discord_pterodactyl_connector
```
- The Windows executable `pterodactylControlBot.exe` is included in this repository for convenience.
- Or build from source:
  ```sh
  go build -o pterodactylControlBot.exe main.go
  ```
  ----
  #### Download
  

### Configuration
1. Copy `config.yml.example` to `config.yml`:
   ```powershell
   Copy-Item config.yml.example config.yml
   ```
2. Edit `config.yml` with your Discord bot token and Pterodactyl API credentials.

## 🛠 Usage
```sh
./pterodactylControlBot.exe
```
Invite the bot to your Discord server and use the available commands to control your servers.

## ⚙️ Example config.yml
```yaml
# This is a sample configuration file for a Discord bot.
# Make sure to replace the placeholders with your actual values.

# The token for your Discord bot. You can get this from the Discord Developer Portal.
discord_token: ""

# Command prefix for the bot. This is what users will type before commands.
command_prefix: ""
# The API token for your Pterodactyl panel. You can generate this in the Pterodactyl admin panel.
api_token: ""
# The URL of your Pterodactyl panel. Make sure to include the protocol (http or https) and trialing slash /.
pterodactyl_url: "https://control.heavynode.com/api/client/"
# The default server ID for the bot to use. This should be the ID of the server you want to manage.
server_id: ""

```

## 🤝 Contributing
Pull requests and issues are welcome!

## 📄 License
MIT
