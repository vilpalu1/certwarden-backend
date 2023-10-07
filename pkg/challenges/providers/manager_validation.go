package providers

import (
	"errors"
	"fmt"
	"legocerthub-backend/pkg/validation"
)

// unsafeValidateDomains verifies that the domains from cfg are all valid
// and also that they're available in manager. p is optional and if specified
// domains will also be condidered valid if they're not available but are
// currently assigned to p.  If validation succeeds, nil is returned, if it
// fails, an error is returned.
func (mgr *Manager) unsafeValidateDomains(cfg providerConfig, p *provider) error {
	// verify every domain ir properly formatted, or verify this is wildcard cfg (* only)
	// and also verify all domains are available in manager
	domains := cfg.Domains()

	// if there are none, invalid
	if len(domains) <= 0 {
		return errors.New("provider config doesn't have any domains (must have at least 1)")
	}

	// validate domain names
	for _, domain := range domains {
		// check validity -or- wildcard
		if !validation.DomainValid(domain, false) && !(len(domains) == 1 && domains[0] == "*") {
			if domain == "*" {
				return errors.New("when using wildcard domain * it must be the only specified domain on the provider")
			}
			return fmt.Errorf("domain %s is not a validly formatted domain", domain)
		}

		// check manager availability
		currentP, exists := mgr.dP[domain]
		if exists && (p == nil || p != currentP) {
			return fmt.Errorf("failed to configure domain %s, each domain can only be configured once", domain)
		}
	}
	return nil
}