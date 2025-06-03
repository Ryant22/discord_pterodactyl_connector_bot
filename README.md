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
# config.yml
# Fill in your details

discord_token: "YOUR_DISCORD_BOT_TOKEN"
pterodactyl:
  api_key: "YOUR_PTERODACTYL_API_KEY"
  panel_url: "https://panel.example.com"
  servers:
    - id: "server-uuid-1"
      name: "Minecraft Server"
    - id: "server-uuid-2"
      name: "CS:GO Server"
```

## 🤝 Contributing
Pull requests and issues are welcome!

## 📄 License
MIT
