## Adze IRC notifier plugin

Plugin for [IRC-Bot](https://github.com/greboid/irc-bot)

Receives notification webhooks from [Adze](https://github.com/greboid/adze) and outputs them to a channel.

#### Configuration

| Flag | Env variable | Default | Description |
|------|-------------|---------|-------------|
| `-rpc-host` | `RPC_HOST` | `localhost` | gRPC server to connect to |
| `-rpc-port` | `RPC_PORT` | `8001` | gRPC server port |
| `-rpc-token` | `RPC_TOKEN` | | gRPC authentication token |
| `-channel` | `CHANNEL` | | Channel to send messages to |
| `-webhook-secret` | `WEBHOOK_SECRET` | | Secret for verifying webhook signatures |
| `-message-prefix` | `MESSAGE_PREFIX` | | Optional prefix to add to the start of each message |
| `-use-dir` | `USE_DIR` | `false` | Use the directory name instead of the project name in messages |

Once configured, set the webhook URL in Adze to `<Bot URL>/adze`, for example:

```
adze -webhook-url https://mybothost.example/adze -webhook-secret <shared-secret>
```

#### Example running

```shell
irc-adze -rpc-host bot -rpc-token <as configured on the bot> -channel #spam -webhook-secret s3cret
```

#### Example Docker Compose

```yaml
services:
  irc-adze:
    image: ghcr.io/csmith/irc-adze
    environment:
      RPC_HOST: bot
      RPC_TOKEN: <as configured on the bot>
      CHANNEL: "#spam"
      WEBHOOK_SECRET: s3cret
      MESSAGE_PREFIX: "[Adze]"
```

## Provenance

This project was primarily created with Claude Code, but with a strong guiding
hand. It's not "vibe coded", but an LLM was still the primary author of most
lines of code. I believe it meets the same sort of standards I'd aim for with
hand-crafted code, but some slop may slip through. I understand if you
prefer not to use LLM-created software, and welcome human-authored alternatives
(I just don't personally have the time/motivation to do so).
