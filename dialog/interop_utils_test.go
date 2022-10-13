// Copyright (c) Microsoft. All rights reserved.
// Licensed under the MIT license. See LICENSE.md file in the project root for full license information.

package dialog

import (
	"testing"

	"github.com/rnhdev2/cognitive-services-speech-sdk-go/common"
)

func TestHandleConversion(t *testing.T) {
	orig := common.SPXHandle(3)
	handle := uintptr2handle(orig)
	dest := handle2uintptr(handle)
	if orig != dest {
		t.Error("Values are not equal")
	}
	if uintptr2handle(dest) != handle {
		t.Error("Values are not equal")
	}
}
