# push-go
A CLI for sending Pushover push notifications to your device.  

## Usage
A config file should exist at `$HOME/.config/.push-go.yaml` containing `PUSHOVER_KEY` and `PUSHOVER_RECIPIENT`. ENV overrides respected.
To use execute with `push-go foo bar`, and a push notification of "foo bar" will be sent to your device.

## Sample Config
if you're supremely lazy, here's the layout
```yaml
PUSHOVER_KEY: KEY
PUSHOVER_RECIPIENT: RECIPIENT_ID
```