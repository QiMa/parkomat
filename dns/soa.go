package dns

import (
	"fmt"
	"strings"

	mdns "github.com/miekg/dns"
	"luckyzune.com/yqbit/parkomat/config"
)

type soaHandler struct {
	name string
}

func NewSOAHandler(name string) *soaHandler {
	return &soaHandler{
		name: name,
	}
}

// Handle produces reply for SOA question
func (s *soaHandler) Handle(msg *mdns.Msg, zone *config.Zone, question mdns.Question) (err error) {
	a := strings.Join(
		[]string{
			question.Name,
			"3600",
			"IN",
			"SOA",
			s.name,
			fmt.Sprintf("admin.%s", question.Name),
			fmt.Sprintf("%d", now()),
			"10000",
			"2400",
			"604800",
			"3600",
		}, " ")

	rr, err := mdns.NewRR(a)
	if err == nil {
		msg.Answer = append(msg.Answer, rr)
	}
	return
}
