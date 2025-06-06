# Discord Pterodactyl Connector

A simple Discord bot written in Go that allows you to control a Pterodactyl game server via Discord commands. The bot can start, stop, restart, kill, and check the status of your server using the Pterodactyl API.

## Features
- Start, stop, restart, and kill your Pterodactyl server from Discord
- Check server status directly from Discord
- Easy configuration via YAML file

## Requirements
- Go 1.18 or newer
- A Discord bot token ([How to create a Discord bot](https://discord.com/developers/applications))
- Pterodactyl API credentials (API key, server ID, and panel URL)

## Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/discord_pterodactyl_connector.git
   cd discord_pterodactyl_connector
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

3. **Configure the bot:**
   - Copy the example config file and edit it:
     ```sh
     cp config.yml.example config.yml
     ```
   - Fill in your Discord bot token, Pterodactyl API token, server ID, and panel URL in `config.yml`.

4. **Build the bot:**
   ```sh
   go build -o pterodactylControlBot.exe main.go
   ```

5. **Run the bot:**
   ```sh
   ./pterodactylControlBot.exe
   ```

## Usage
Invite your bot to your Discord server. Use the following commands in any channel the bot can read:

> **Security Note:** Only Discord server administrators can use the bot commands. Non-administrators will be denied access for security reasons.

- `<prefix>start` — Start the server
- `<prefix>stop` — Stop the server
- `<prefix>restart` — Restart the server
- `<prefix>kill` — Kill the server
- `<prefix>status` — Get the server status
- `<prefix>help` — Show help message

Replace `<prefix>` with the command prefix you set in `config.yml` (default is `!`).

## Example `config.yml`
```yaml
discord_token: "YOUR_DISCORD_BOT_TOKEN"
command_prefix: "!"
api_token: "YOUR_PTERODACTYL_API_TOKEN"
pterodactyl_url: "https://panel.example.com/api/client/"
server_id: "YOUR_SERVER_ID"
```

## Notes
- Make sure your Pterodactyl API key has the correct permissions for server power actions.
- The bot must be running for commands to work.
- For production use, consider running the bot as a service or in a screen/tmux session.

## License
MIT License
