# Transak Webhook Data Decoder

Decoding Transak Order Webhook Data in GO!

This will decode the incoming webhook data and output it to `decoded_webhook_data.json` in the project directory.

## Dependencies

- `github.com/golang-jwt/jwt/v5`: For decoding JWTs in webhook payloads.
