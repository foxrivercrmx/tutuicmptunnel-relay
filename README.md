# tutuicmptunnel Relay

## Installation
* Create workers on your Cloudflare account
* Copy-paste workers.js code to your cloudflare workers, and deploy
* Build backend relay
```bash
git clone https://github.com/foxrivercrmx/tutuicmptunnel-relay
cd tutuicmptunnel-relay
CGO_ENABLED=0 go build -ldflags="-s -w" -o tutu-relay main.go
```
* *optional:*
```bash
sudo mv tutu-relay /usr/local/bin/
```
## Usage
*example:*
```bash
tutu-relay -path=/your/path -psk=strong-password -mode=tuctl -listen=:8080
```

### Notes
* Work best with [tutuicmptunnel-gui](https://github.com/foxrivercrmx/tutuicmptunnel-gui)
* Made with ❤️ and ☕