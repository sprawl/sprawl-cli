# Sprawl-cli
Sprawl-cli is a command line client for using Sprawl nodes directly. It's based on https://github.com/fiorix/protoc-gen-cobra.

## Usage

Since Sprawl uses port 1337 by default, the commands need the flag -s "address":"port". If you get any response, the commands should work.

Joining channels:
```bash
echo -ne '{"Asset": "BTC", "Counterasset": "ETH"}\n' | ./sprawl-cli channelhandler join -s localhost:1337
```

Creating orders:
```bash
echo -ne '{"ChannelID": "QlRDLEVUSA==", "Asset":"ETH", "CounterAsset":"BTC", "Amount":52167, "Price":0.8}\n' | ./sprawl-cli orderhandler create          
-s localhost:1337
```

Fetching all saved orders:
```bash
echo -ne '{}\n' | ./sprawl-cli orderhandler getallorders -s localhost:1337
```
