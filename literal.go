// Copyright 2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package scanner

import "strconv"

func ParseInt(format int, literal string) (int64, error) {
	return strconv.ParseInt(literal, format, 64)
}

func ParseUint(format int, literal string) (uint64, error) {
	return strconv.ParseUint(literal, format, 64)
}
