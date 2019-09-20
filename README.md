# Sprawl-cli
Sprawl-cli is a command line client for using [Sprawl](https://github.com/eqlabs/sprawl) nodes directly. It uses generated protobuf code from Sprawl to create a protobuf client that's usable as a CLI.

## Usage

The commands need the flag -s "address":"port". If you get any response, the commands should work.

Build the tool:
```bash
go build
```

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
