package main

import ()

var TwilioAccountSid string = ""

type TwilioNumber struct {
	AddressRequirements string `json:"address_requirements"`
	Beta                bool   `json:"beta"`
	Capabilities        struct {
		MMS   bool `json:"MMS"`
		SMS   bool `json:"SMS"`
		Voice bool `json:"voice"`
	} `json:"capabilities"`
	FriendlyName string `json:"friendly_name"`
	IsoCountry   string `json:"iso_country"`
	Lata         string `json:"lata"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	PhoneNumber  string `json:"phone_number"`
	PostalCode   string `json:"postal_code"`
	RateCenter   string `json:"rate_center"`
	Region       string `json:"region"`
}

type TwilioAvailableNumbers struct {
	Numbers []TwilioNumber `json:"available_phone_numbers"`
}

type TwilioAddedNumber struct {
	AccountSid   string `json:"account_sid"`
	APIVersion   string `json:"api_version"`
	Capabilities struct {
		Mms   bool `json:"mms"`
		Sms   bool `json:"sms"`
		Voice bool `json:"voice"`
	} `json:"capabilities"`
	DateCreated          string `json:"date_created"`
	DateUpdated          string `json:"date_updated"`
	FriendlyName         string `json:"friendly_name"`
	PhoneNumber          string `json:"phone_number"`
	Sid                  string `json:"sid"`
	SmsApplicationSid    string `json:"sms_application_sid"`
	SmsFallbackMethod    string `json:"sms_fallback_method"`
	SmsFallbackURL       string `json:"sms_fallback_url"`
	SmsMethod            string `json:"sms_method"`
	SmsURL               string `json:"sms_url"`
	StatusCallback       string `json:"status_callback"`
	StatusCallbackMethod string `json:"status_callback_method"`
	URI                  string `json:"uri"`
	VoiceApplicationSid  string `json:"voice_application_sid"`
	VoiceCallerIDLookup  string `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod  string `json:"voice_fallback_method"`
	VoiceFallbackURL     string `json:"voice_fallback_url"`
	VoiceMethod          string `json:"voice_method"`
	VoiceURL             string `json:"voice_url"`
}

func TwilioInit(appSid string) {

}

func TwillioGetAvailableNumbers() {

}

func TwilioGetNumber(number string, description string) {

}

func TwilioDeleteNumber() {

}
