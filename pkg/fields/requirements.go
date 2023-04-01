// Copyright (c) 2023 khh403. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fields

import "github.com/khh403/component-base/pkg/selection"

// Requirements is AND of all requirements.
type Requirements []Requirement

// Requirement contains a field, a value, and an operator that relates the field and value.
// This is currently for reading internal selection information of field selector.
type Requirement struct {
	Operator selection.Operator
	Field    string
	Value    string
}
