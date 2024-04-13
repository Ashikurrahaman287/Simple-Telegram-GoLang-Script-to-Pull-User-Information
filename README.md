# Telegram Username Collector

This Go application is designed to collect usernames from Telegram groups and save them to a text file. It utilizes the Telegram Bot API for communication with Telegram.

## Features

- Connects to Telegram using a bot API token.
- Collects usernames from a specified Telegram group.
- Saves usernames to a text file for further processing.

## Usage

1. Replace the placeholder with your own Telegram API token in the `main` function.
2. Specify the Telegram group from which you want to collect usernames by modifying the `groupName` variable in the `main` function.
3. Run the application.

```bash
go run main.go
```

## Requirements

This application depends on the `github.com/go-telegram-bot-api/telegram-bot-api` package. Install it using the following command:

```bash
go get -u github.com/go-telegram-bot-api/telegram-bot-api
```

## Contact

Feel free to reach out to me for any questions or feedback:

- Email: Ashik@spudblocks.com
- Telegram: @Ashik_Spudblocks
