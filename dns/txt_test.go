package dns

import (
	"testing"

	mdns "github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
	"luckyzune.com/yqbit/parkomat/config"
)

func TestTXTHandle(t *testing.T) {
	msg := &mdns.Msg{}

	zone := &config.Zone{
		TXT: `never gonna
		give you up
`,
	}

	question := mdns.Question{
		Name: "test.com",
	}

	txt := &txtHandler{}

	err := txt.Handle(msg, zone, question)
	assert.Nil(t, err)

	expectedMsg := &mdns.Msg{}
	rr, err := mdns.NewRR("test.com. 3600 IN TXT never gonna")
	assert.Nil(t, err)
	expectedMsg.Answer = append(expectedMsg.Answer, rr)
	rr, err = mdns.NewRR("test.com. 3600 IN TXT give you up")
	assert.Nil(t, err)
	expectedMsg.Answer = append(expectedMsg.Answer, rr)

	assert.Exactly(t, msg.Answer, expectedMsg.Answer)
}
