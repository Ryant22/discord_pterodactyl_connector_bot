# Discord Pterodactyl Connector Bot

A Discord bot to control your Pterodactyl game server from Discord. Supports starting, stopping, restarting, killing, and checking the status of your server via Discord commands.

---

## Features
- Start, stop, restart, kill, and check the status of your Pterodactyl server via Discord commands
- Easy YAML configuration
- Works with any Pterodactyl panel supporting the client API

---

## Setup

### 1. Download and Extract
- Download the project folder or the `discord_pterodactyl_connector.rar` file.
- If you downloaded the `.rar` file, extract it to a folder on your computer.

### 2. Configure the Bot
- Copy `config.yml.example` to `config.yml`.
- Edit `config.yml` and fill in:
  - `discord_token`: Your Discord bot token (from the Discord Developer Portal)
  - `command_prefix`: The prefix for commands (e.g., `!`)
  - `api_token`: Your Pterodactyl API token
  - `pterodactyl_url`: The URL to your Pterodactyl panel’s API (should end with `/api/client/`)
  - `server_id`: The ID of the server you want to control

### 3. Running the Bot

#### Option A: Pre-built Executable
- Double-click `pterodactylControlBot.exe` to start the bot.

#### Option B: From Source
- Install Go (https://go.dev/dl/)
- Open a terminal in the project folder
- Run: `go run main.go`

---

## Usage
- Invite your bot to your Discord server (generate an invite link in the Discord Developer Portal)
- Use commands in any channel the bot can read:
  - `<prefix>start` — Start the server
  - `<prefix>stop` — Stop the server
  - `<prefix>restart` — Restart the server
  - `<prefix>kill` — Kill the server
  - `<prefix>status` — Get server status
  - `<prefix>help` — Show help message

---

## Example `config.yml`
```yaml
discord_token: "YOUR_DISCORD_BOT_TOKEN"
command_prefix: "!"
api_token: "YOUR_PTERODACTYL_API_TOKEN"
pterodactyl_url: "https://your.panel.url/api/client/"
server_id: "YOUR_SERVER_ID"
```

---

## Troubleshooting
- Double-check your `config.yml` values for typos
- Ensure your tokens and URLs are correct
- If you see errors, check the terminal or log output for details

---

## Contributing
Pull requests and issues are welcome! Please open an issue for bugs or feature requests.

---

## License
MIT License
