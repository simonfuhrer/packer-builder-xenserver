package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
<<<<<<< HEAD
	xsclient "github.com/simonfuhrer/go-xenserver-client"
=======
	xsclient "github.com/xenserver/go-xenserver-client"
>>>>>>> aa0bbcae25c2db138b23c8f008f5948721a18cfc
)

type StepBootWait struct{}

func (self *StepBootWait) Run(state multistep.StateBag) multistep.StepAction {
	client := state.Get("client").(xsclient.XenAPIClient)
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)

	instance, _ := client.GetVMByUuid(state.Get("instance_uuid").(string))
	ui.Say("Unpausing VM " + state.Get("instance_uuid").(string))
	instance.Unpause()

	if int64(config.BootWait) > 0 {
		ui.Say(fmt.Sprintf("Waiting %s for boot...", config.BootWait))
		err := InterruptibleWait{Timeout: config.BootWait}.Wait(state)
		if err != nil {
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}
	return multistep.ActionContinue
}

func (self *StepBootWait) Cleanup(state multistep.StateBag) {}
