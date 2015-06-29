package main

import (
	"github.com/mitchellh/packer/packer/plugin"
<<<<<<< HEAD
	"github.com/simonfuhrer/packer-builder-xenserver/builder/xenserver/xva"
=======
	"github.com/rdobson/packer-builder-xenserver/builder/xenserver/xva"
>>>>>>> aa0bbcae25c2db138b23c8f008f5948721a18cfc
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(xva.Builder))
	server.Serve()
}
