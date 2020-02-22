/*
Package butterfly : The library for butterfly

Copyright(c) 2020 Star Inc. All Rights Reserved.
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package butterfly

import "fmt"

// DeBug :
func DeBug(where string, err error) bool {
	if err != nil {
		fmt.Printf("NepCore Error #%s\nReason:\n%s\n\n", where, err)
		return false
	}
	return true
}
