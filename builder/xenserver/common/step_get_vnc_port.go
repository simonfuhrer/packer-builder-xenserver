package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
        xsclient "github.com/simonfuhrer/go-xenserver-client"

	"strconv"
)

type StepGetVNCPort struct{}

func (self *StepGetVNCPort) Run(state multistep.StateBag) multistep.StepAction {
        client := state.Get("client").(xsclient.XenAPIClient)
        ui := state.Get("ui").(packer.Ui)

        ui.Say("Step: Get VNC Console")

        uuid := state.Get("instance_uuid").(string)
        instance, err := client.GetVMByUuid(uuid)
        if err != nil {
                ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
                return multistep.ActionHalt
        }
        consoles, err := instance.GetConsoles()
        if err != nil {
                ui.Error(fmt.Sprintf("Unable to get VM Console from UUID '%s': %s", uuid, err.Error()))
                return multistep.ActionHalt
        }
        var consoleLocation string
        for _, con := range consoles {
                conrec, _ := con.GetRecord()
                if conrec["protocol"] == "rfb" && conrec["VM"] != "" {
			consoleLocation = conrec["location"].(string)
			break
                }
        }
	if consoleLocation == "" {
		ui.Error(fmt.Sprintf("Unable to get VM Console Location '%s'", uuid, err.Error()))
		return multistep.ActionHalt
        }
	ui.Say("Step: STORE VNC LOCATION")
	state.Put("instance_vnc_location", )

	return multistep.ActionContinue
}

func (self *StepGetVNCPort) Cleanup(state multistep.StateBag) {
}

func InstanceVNCPort(state multistep.StateBag) (uint, error) {
	vncPort := state.Get("instance_vnc_port").(uint)
	return vncPort, nil
}

func InstanceVNCIP(state multistep.StateBag) (string, error) {
	// The port is in Dom0, so we want to forward from localhost
	return "127.0.0.1", nil
}
