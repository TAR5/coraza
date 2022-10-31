// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package operators

import (
	"github.com/corazawaf/coraza/v3/rules"
)

func newGeoLookup(rules.OperatorOptions) (rules.Operator, error) {
	return &unconditionalMatch{}, nil
}
