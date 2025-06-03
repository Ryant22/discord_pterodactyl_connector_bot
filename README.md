# Discord Pterodactyl Connector Bot

A simple Discord bot to control your Pterodactyl game server from Discord.

---

## Features
- Start, stop, restart, kill, and check the status of your Pterodactyl server via Discord commands.
- Easy configuration with a YAML file.

---

## Setup Instructions (For Non-Technical Users)

### 1. Download and Extract
- Download the project folder or the `discord_pterodactyl_connector.rar` file.
- If you downloaded the `.rar` file, extract it to a folder on your computer.

### 2. Configure the Bot
- Find the file named `config.yml.example` in the folder.
- Make a copy of it and rename the copy to `config.yml`.
- Open `config.yml` with Notepad or any text editor.
- Fill in the following fields:
  - `discord_token`: Your Discord bot token (from the Discord Developer Portal).
  - `command_prefix`: The symbol or word you want to use before commands (e.g., `!`).
  - `api_token`: Your Pterodactyl API token (from your Pterodactyl panel).
  - `pterodactyl_url`: The URL to your Pterodactyl panel’s API (should end with `/api/client/`).
  - `server_id`: The ID of the server you want to control.

### 3. Run the Bot

#### Option A: Using the Pre-built EXE
- Double-click `pterodactylControlBot.exe` to start the bot.

#### Option B: Running from Source (Advanced)
- Install Go from https://go.dev/dl/
- Open a command prompt in the project folder.
- Run: `go run main.go`
- The bot will start and show a message when it’s running.

### 4. Invite the Bot to Your Discord Server
- Use the Discord Developer Portal to generate an invite link for your bot.
- Make sure the bot has permission to read and send messages.

### 5. Use the Bot
- In your Discord server, type your chosen prefix and a command, for example: `!start`, `!stop`, `!status`, `!restart`, `!kill`, or `!help`.

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
- Double-check your `config.yml` values for typos.
- Make sure your tokens and URLs are correct.
- If you see errors, check the command prompt or log output for details.

If you need more help, contact the project maintainer.
