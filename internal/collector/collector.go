package collector

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"

	"github.com/pkg/errors"
)

// Collector for the stacklight working durations
type Collector struct {
	mu       sync.RWMutex
	duration map[string]time.Duration
	state    map[string]bool
	labels   map[string]string
}

func New(colors []string, stateFile string, labels map[string]string) (*Collector, error) {
	// read previus state
	rs, err := ioutil.ReadFile(stateFile)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't read the state file")
	}
	err = json.Unmarshal(rs)
}
