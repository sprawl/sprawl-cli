package main

import (
	"github.com/eqlabs/sprawl-cli/cmd"
	"github.com/eqlabs/sprawl-cli/pb"
)

func init() {
	cmd.RootCmd.AddCommand(pb.OrderHandlerClientCommand)
	cmd.RootCmd.AddCommand(pb.ChannelHandlerClientCommand)
}
