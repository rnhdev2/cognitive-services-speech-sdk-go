// Copyright (c) Microsoft. All rights reserved.
// Licensed under the MIT license. See LICENSE.md file in the project root for full license information.

package speaker

import (
	"github.com/rnhdev2/cognitive-services-speech-sdk-go/common"
)

// #include <stdlib.h>
// #include <speechapi_c_result.h>
//
import "C"

// VoiceProfileCancellationDetails contains detailed information about why a result was canceled.
// Added in version 1.21.0
type VoiceProfileCancellationDetails struct {
	Reason       common.CancellationReason
	ErrorCode    common.CancellationErrorCode
	ErrorDetails string
}

// NewCancellationDetailsFromVoiceProfileResult creates the object from the speech synthesis result.
func NewCancellationDetailsFromVoiceProfileResult(result *VoiceProfileResult) (*VoiceProfileCancellationDetails, error) {
	cancellationDetails := new(VoiceProfileCancellationDetails)
	/* Reason */
	var cReason C.Result_CancellationReason
	ret := uintptr(C.result_get_reason_canceled(result.handle, &cReason))
	if ret != C.SPX_NOERROR {
		return nil, common.NewCarbonError(ret)
	}
	cancellationDetails.Reason = (common.CancellationReason)(cReason)
	/* ErrorCode */
	var cCode C.Result_CancellationErrorCode
	ret = uintptr(C.result_get_canceled_error_code(result.handle, &cCode))
	if ret != C.SPX_NOERROR {
		return nil, common.NewCarbonError(ret)
	}
	cancellationDetails.ErrorCode = (common.CancellationErrorCode)(cCode)
	cancellationDetails.ErrorDetails = result.Properties.GetProperty(common.CancellationDetailsReasonDetailedText, "")
	return cancellationDetails, nil
}
