// Copyright (c) Microsoft. All rights reserved.
// Licensed under the MIT license. See LICENSE.md file in the project root for full license information.

package speaker_recognition

import (
	"fmt"
	"time"

	"github.com/rnhdev2/cognitive-services-speech-sdk-go/audio"
	"github.com/rnhdev2/cognitive-services-speech-sdk-go/common"
	"github.com/rnhdev2/cognitive-services-speech-sdk-go/speaker"
	"github.com/rnhdev2/cognitive-services-speech-sdk-go/speech"
)

func GetNewVoiceProfileFromClient(client *speaker.VoiceProfileClient, expectedType common.VoiceProfileType) *speaker.VoiceProfile {
	future := client.CreateProfileAsync(expectedType, "en-US")
	outcome := <-future
	if outcome.Failed() {
		fmt.Println("Got an error creating profile: ", outcome.Error.Error())
		return nil
	}
	profile := outcome.Profile
	_, err := profile.Id()
	if err != nil {
		fmt.Println("Unexpected error creating profile id: ", err)
		return nil
	}
	profileType, err := profile.Type()
	if err != nil {
		fmt.Println("Unexpected error getting profile type: ", err)
		return nil
	}
	if profileType != expectedType {
		fmt.Println("Profile type does not match expected type")
		return nil
	}
	return profile
}

func EnrollProfile(client *speaker.VoiceProfileClient, profile *speaker.VoiceProfile, audioConfig *audio.AudioConfig) {
	enrollmentReason, currentReason := common.EnrollingVoiceProfile, common.EnrollingVoiceProfile
	var currentResult *speaker.VoiceProfileEnrollmentResult
	expectedEnrollmentCount := 1
	for currentReason == enrollmentReason {
		enrollFuture := client.EnrollProfileAsync(profile, audioConfig)
		enrollOutcome := <-enrollFuture
		if enrollOutcome.Failed() {
			fmt.Println("Got an error enrolling profile: ", enrollOutcome.Error.Error())
			return
		}
		currentResult = enrollOutcome.Result
		currentReason = currentResult.Reason
		if currentResult.EnrollmentsCount != expectedEnrollmentCount {
			fmt.Println("Unexpected enrollments for profile: ", currentResult.RemainingEnrollmentsCount)
		}
		expectedEnrollmentCount += 1
	}
	if currentReason != common.EnrolledVoiceProfile {
		fmt.Println("Unexpected result enrolling profile: ", currentResult)
	}
}

func DeleteProfile(client *speaker.VoiceProfileClient, profile *speaker.VoiceProfile) {
	deleteFuture := client.DeleteProfileAsync(profile)
	deleteOutcome := <-deleteFuture
	if deleteOutcome.Failed() {
		fmt.Println("Got an error deleting profile: ", deleteOutcome.Error.Error())
		return
	}
	result := deleteOutcome.Result
	if result.Reason != common.DeletedVoiceProfile {
		fmt.Println("Unexpected result deleting profile: ", result)
	}
}

func IndependentIdentification(subscription string, region string, file string) {
	config, err := speech.NewSpeechConfigFromSubscription(subscription, region)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return
	}
	defer config.Close()
	client, err := speaker.NewVoiceProfileClientFromConfig(config)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return
	}
	defer client.Close()
	audioConfig, err := audio.NewAudioConfigFromWavFileInput(file)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return
	}
	defer audioConfig.Close()
	<-time.After(10 * time.Second)
	expectedType := common.VoiceProfileType(1)

	profile := GetNewVoiceProfileFromClient(client, expectedType)
	if profile == nil {
		fmt.Println("Error creating profile")
		return
	}
	defer profile.Close()

	EnrollProfile(client, profile, audioConfig)

	profiles := []*speaker.VoiceProfile{profile}
	model, err := speaker.NewSpeakerIdentificationModelFromProfiles(profiles)
	if err != nil {
		fmt.Println("Error creating Identification model: ", err)
	}
	if model == nil {
		fmt.Println("Error creating Identification model: nil model")
		return
	}
	identifyAudioConfig, err := audio.NewAudioConfigFromWavFileInput(file)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return
	}
	defer identifyAudioConfig.Close()
	speakerRecognizer, err := speaker.NewSpeakerRecognizerFromConfig(config, identifyAudioConfig)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return
	}
	identifyFuture := speakerRecognizer.IdentifyOnceAsync(model)
	identifyOutcome := <-identifyFuture
	if identifyOutcome.Failed() {
		fmt.Println("Got an error identifying profile: ", identifyOutcome.Error.Error())
		return
	}
	identifyResult := identifyOutcome.Result
	if identifyResult.Reason != common.RecognizedSpeakers {
		fmt.Println("Got an unexpected result identifying profile: ", identifyResult)
	}
	expectedID, _ := profile.Id()
	if identifyResult.ProfileID != expectedID {
		fmt.Println("Got an unexpected profile id identifying profile: ", identifyResult.ProfileID)
	}
	if identifyResult.Score < 1.0 {
		fmt.Println("Got an unexpected score identifying profile: ", identifyResult.Score)
	}

	DeleteProfile(client, profile)
}
