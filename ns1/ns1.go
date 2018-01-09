package ns1

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/xenolf/lego/acme"
	"github.com/xenolf/lego/providers/dns/ns1"
)

func init() {
	caddytls.RegisterDNSProvider("ns1", NewDNSProvider)
}

// NewDNSProvider returns a new ns1.DNSProvider DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(1): credentials[0] = API key

func NewDNSProvider(credentials ...string) (acme.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return ns1.NewDNSProvider()
	case 1:
		return ns1.NewDNSProviderCredentials(credentials[0])
	default:
		return nil, errors.New("invalid credentials length")
	}
}
