package dns

import (
	"strings"

	mdns "github.com/miekg/dns"
	"luckyzune.com/yqbit/parkomat/config"
)

type txtHandler struct {
}

// Handle produces reply for TXT question
func (t *txtHandler) Handle(msg *mdns.Msg, zone *config.Zone, question mdns.Question) (err error) {
	for _, txt := range strings.Split(zone.TXT, "\n") {
		txt = strings.Trim(txt, " ")
		if txt != "" {
			s := strings.Join([]string{
				question.Name,
				"3600",
				"IN",
				"TXT",
				txt,
			}, " ")

			rr, err := mdns.NewRR(s)
			if err == nil {
				msg.Answer = append(msg.Answer, rr)
			}
		}
	}
	return
}
