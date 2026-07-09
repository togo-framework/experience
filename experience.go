// Package experience is the Experience content pack for togo: Experience timeline resource + a CV view.
//
// A content pack bundles resource definitions (model + sqlc + Atlas + REST +
// GraphQL + a page) that a togo app pulls in. Generate the resource with
// `togo make:resource` in the host app; this plugin wires the provider hook.
package experience

import "github.com/togo-framework/togo"

func init() {
	togo.RegisterProviderFunc("experience", togo.PriorityService, func(k *togo.Kernel) error {
		return nil
	})
}
