package sakuracloud

import (
	"context"
	"testing"
	"time"

	"github.com/hashicorp/packer/helper/multistep"
	"github.com/stretchr/testify/assert"
)

func TestStepBootWait(t *testing.T) {
	ctx := context.Background()
	step := &stepBootWait{}
	conf := dummyConfig()
	conf.BootWait = 10 * time.Millisecond

	state := dummyMinimumStateBag(&conf)

	// run
	start := time.Now()
	action := step.Run(ctx, state)
	end := time.Now()

	expect := conf.BootWait
	actual := end.Sub(start)

	assert.True(t, actual.Seconds() > expect.Seconds())
	assert.Equal(t, multistep.ActionContinue, action)
}
