package dns

import (
	"strings"

	mdns "github.com/miekg/dns"
	"luckyzune.com/yqbit/parkomat/config"
)

type aHandler struct {
}

// Handle produces reply for A question
func (a *aHandler) Handle(msg *mdns.Msg, zone *config.Zone, question mdns.Question) (err error) {
	s := strings.Join(
		[]string{
			question.Name,
			"3600",
			"IN",
			"A",
			zone.A,
		}, " ")

	rr, err := mdns.NewRR(s)
	if err == nil {
		msg.Answer = append(msg.Answer, rr)
	}
	return
}
