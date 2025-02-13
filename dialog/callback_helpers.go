// Copyright (c) Microsoft. All rights reserved.
// Licensed under the MIT license. See LICENSE.md file in the project root for full license information.

package dialog

import (
	"sync"

	"github.com/rnhdev2/cognitive-services-speech-sdk-go/speech"
)

// #include <speechapi_c_common.h>
// #include <speechapi_c_recognizer.h>
// #include <speechapi_c_dialog_service_connector.h>
import "C"

var mu sync.Mutex
var sessionStartedCallbacks = make(map[C.SPXHANDLE]speech.SessionEventHandler)

func registerSessionStartedCallback(handler speech.SessionEventHandler, handle C.SPXHANDLE) {
	mu.Lock()
	defer mu.Unlock()
	sessionStartedCallbacks[handle] = handler
}

func getSessionStartedCallback(handle C.SPXHANDLE) speech.SessionEventHandler {
	mu.Lock()
	defer mu.Unlock()
	return sessionStartedCallbacks[handle]
}

//export dialogFireEventSessionStarted
func dialogFireEventSessionStarted(handle C.SPXRECOHANDLE, eventHandle C.SPXEVENTHANDLE) {
	handler := getSessionStartedCallback(handle)
	event, err := speech.NewSessionEventArgsFromHandle(handle2uintptr(eventHandle))
	if err != nil || handler == nil {
		C.recognizer_event_handle_release(handle)
		return
	}
	handler(*event)
}

var sessionStoppedCallbacks = make(map[C.SPXHANDLE]speech.SessionEventHandler)

func registerSessionStoppedCallback(handler speech.SessionEventHandler, handle C.SPXHANDLE) {
	mu.Lock()
	defer mu.Unlock()
	sessionStoppedCallbacks[handle] = handler
}

func getSessionStoppedCallback(handle C.SPXHANDLE) speech.SessionEventHandler {
	mu.Lock()
	defer mu.Unlock()
	return sessionStoppedCallbacks[handle]
}

//export dialogFireEventSessionStopped
func dialogFireEventSessionStopped(handle C.SPXRECOHANDLE, eventHandle C.SPXEVENTHANDLE) {
	handler := getSessionStoppedCallback(handle)
	event, err := speech.NewSessionEventArgsFromHandle(handle2uintptr(eventHandle))
	if err != nil || handler == nil {
		C.recognizer_event_handle_release(handle)
		return
	}
	handler(*event)
}

var recognizedCallbacks = make(map[C.SPXHANDLE]speech.SpeechRecognitionEventHandler)

func registerRecognizedCallback(handler speech.SpeechRecognitionEventHandler, handle C.SPXHANDLE) {
	mu.Lock()
	defer mu.Unlock()
	recognizedCallbacks[handle] = handler
}

func getRecognizedCallback(handle C.SPXHANDLE) speech.SpeechRecognitionEventHandler {
	mu.Lock()
	defer mu.Unlock()
	return recognizedCallbacks[handle]
}

//export dialogFireEventRecognized
func dialogFireEventRecognized(handle C.SPXRECOHANDLE, eventHandle C.SPXEVENTHANDLE) {
	handler := getRecognizedCallback(handle)
	event, err := speech.NewSpeechRecognitionEventArgsFromHandle(handle2uintptr(eventHandle))
	if err != nil || handler == nil {
		C.recognizer_event_handle_release(handle)
		return
	}
	handler(*event)
}

var recognizingCallbacks = make(map[C.SPXHANDLE]speech.SpeechRecognitionEventHandler)

func registerRecognizingCallback(handler speech.SpeechRecognitionEventHandler, handle C.SPXHANDLE) {
	mu.Lock()
	defer mu.Unlock()
	recognizingCallbacks[handle] = handler
}

func getRecognizingCallback(handle C.SPXHANDLE) speech.SpeechRecognitionEventHandler {
	mu.Lock()
	defer mu.Unlock()
	return recognizingCallbacks[handle]
}

//export dialogFireEventRecognizing
func dialogFireEventRecognizing(handle C.SPXRECOHANDLE, eventHandle C.SPXEVENTHANDLE) {
	handler := getRecognizingCallback(handle)
	event, err := speech.NewSpeechRecognitionEventArgsFromHandle(handle2uintptr(eventHandle))
	if err != nil || handler == nil {
		C.recognizer_event_handle_release(handle)
		return
	}
	handler(*event)
}

var canceledCallbacks = make(map[C.SPXHANDLE]speech.SpeechRecognitionCanceledEventHandler)

func registerCanceledCallback(handler speech.SpeechRecognitionCanceledEventHandler, handle C.SPXHANDLE) {
	mu.Lock()
	defer mu.Unlock()
	canceledCallbacks[handle] = handler
}

func getCanceledCallback(handle C.SPXHANDLE) speech.SpeechRecognitionCanceledEventHandler {
	mu.Lock()
	defer mu.Unlock()
	return canceledCallbacks[handle]
}

//export dialogFireEventCanceled
func dialogFireEventCanceled(handle C.SPXRECOHANDLE, eventHandle C.SPXEVENTHANDLE) {
	handler := getCanceledCallback(handle)
	event, err := speech.NewSpeechRecognitionCanceledEventArgsFromHandle(handle2uintptr(eventHandle))
	if err != nil || handler == nil {
		C.recognizer_event_handle_release(handle)
		return
	}
	handler(*event)
}

var activityReceivedCallbacks = make(map[C.SPXHANDLE]ActivityReceivedEventHandler)

func registerActivityReceivedCallback(handler ActivityReceivedEventHandler, handle C.SPXHANDLE) {
	mu.Lock()
	defer mu.Unlock()
	activityReceivedCallbacks[handle] = handler
}

func getActivityReceivedCallback(handle C.SPXHANDLE) ActivityReceivedEventHandler {
	mu.Lock()
	defer mu.Unlock()
	return activityReceivedCallbacks[handle]
}

//export dialogFireEventActivityReceived
func dialogFireEventActivityReceived(handle C.SPXRECOHANDLE, eventHandle C.SPXEVENTHANDLE) {
	handler := getActivityReceivedCallback(handle)
	event, err := NewActivityReceivedEventArgsFromHandle(handle2uintptr(eventHandle))
	if err != nil || handler == nil {
		C.dialog_service_connector_activity_received_event_release(handle)
		return
	}
	handler(*event)
}
