package sakuracloud

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
	"github.com/sacloud/packer-builder-sakuracloud/iaas"
)

type stepPrepareISO struct {
	Debug bool
}

func (s *stepPrepareISO) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	isoImageClient := state.Get("isoImageClient").(iaas.ISOImageClient)
	config := state.Get("config").(Config)
	ui := state.Get("ui").(packer.Ui)

	stepStartMsg(ui, s.Debug, "PrepareISO")

	isoID := config.ISOImageID
	if isoID == 0 {
		isoID = state.Get("iso_id").(int64)
	}

	config.ISOImageID = isoID

	iso, err := isoImageClient.Read(isoID)
	if err != nil {
		err := fmt.Errorf("Error invalid ISO image ID: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	if !iso.IsAvailable() {
		err := fmt.Errorf("Error invalid ISO image Status: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	state.Put("config", config)
	stepEndMsg(ui, s.Debug, "PrepareISO")
	return multistep.ActionContinue
}

func (s *stepPrepareISO) Cleanup(state multistep.StateBag) {
}
