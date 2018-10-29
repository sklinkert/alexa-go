package alexa

// constants

// built in intents
const (
	//HelpIntent is the Alexa built-in Help Intent
	HelpIntent = "AMAZON.HelpIntent"

	//CancelIntent is the Alexa built-in Cancel Intent
	CancelIntent = "AMAZON.CancelIntent"

	//StopIntent is the Alexa built-in Stop Intent
	StopIntent = "AMAZON.StopIntent"

	// PauseIntent - Built-in for pause.
	PauseIntent = "AMAZON.PauseIntent"
	// ResumeIntent - Built-in for resume.
	ResumeIntent = "AMAZON.ResumeIntent"
	// NextIntent - Built-in for next (e.g. next audio track).
	NextIntent = "AMAZON.NextIntent"
	// PreviousIntent - Built-in for previous (e.g. previous audio track).
	PreviousIntent = "AMAZON.PreviousIntent"
	// RepeatIntent - Built-in for repeat (e.g. repeat current audio track).
	RepeatIntent = "AMAZON.RepeatIntent"
	// ShuffleOnIntent - Built-in for enabling shuffle mode
	ShuffleOnIntent = "AMAZON.ShuffleOnIntent"
	// ShuffleOffIntent - Built-in for disabling shuffle mode
	ShuffleOffIntent = "AMAZON.ShuffleOffIntent"
	// LoopOnIntent - Built-in for enabling loop mode
	LoopOnIntent = "AMAZON.LoopOnIntent"
	// LoopOffIntent - Built-in for disabling loop mode
	LoopOffIntent = "AMAZON.LoopOffIntent"
	// StartOverIntent - Built-in for starting over
	StartOverIntent = "AMAZON.StartOverIntent"
	// NoIntent - Built-in for "no" response
	NoIntent = "AMAZON.NoIntent"
	// YesIntent - Built-in for "yes" response
	YesIntent = "AMAZON.YesIntent"

	// LaunchRequest - Built-in intent for LaunchRequest
	LaunchRequest = "LaunchRequest"
	// SystemExceptionEncountered - Built-in intent for unexpected exceptions.
	SystemExceptionEncountered = "System.ExceptionEncountered"
	// SessionEndedRequest - Built-in intent for the end of skill session
	SessionEndedRequest = "SessionEndedRequest"

	// AudioPlayerPlaybackStarted - Playback Started
	AudioPlayerPlaybackStarted = "AudioPlayer.PlaybackStarted"
	// AudioPlayerPlaybackNearlyFinished - Playback Nearly Finished
	AudioPlayerPlaybackNearlyFinished = "AudioPlayer.PlaybackNearlyFinished"
	// AudioPlayerPlaybackFinished - Playback Finished
	AudioPlayerPlaybackFinished = "AudioPlayer.PlaybackFinished"
	// AudioPlayerPlaybackStopped - Playback Stopped (by user)
	AudioPlayerPlaybackStopped = "AudioPlayer.PlaybackStopped"
	// AudioPlayerPlaybackFailed - Playback Failed
	AudioPlayerPlaybackFailed = "AudioPlayer.PlaybackFailed"

	// PlaybackControllerPlayCommandIssued - Play command issued by user
	PlaybackControllerPlayCommandIssued = "PlaybackController.PlayCommandIssued"
	// PlaybackControllerPreviousCommandIssued - Previous command issued by user
	PlaybackControllerPreviousCommandIssued = "PlaybackController.PreviousCommandIssued"
	// PlaybackControllerNextCommandIssued - Next command issued by user
	PlaybackControllerNextCommandIssued = "PlaybackController.NextCommandIssued"
	// PlaybackControllerPauseCommandIssued - Pause command issued by user
	PlaybackControllerPauseCommandIssued = "PlaybackController.PauseCommandIssued"
)

// locales
const (
	// LocaleItalian is the locale for Italian
	LocaleItalian = "it-IT"

	// LocaleGerman is the locale for standard dialect German
	LocaleGerman = "de-DE"

	// LocaleAustralianEnglish is the locale for Australian English
	LocaleAustralianEnglish = "en-AU"

	//LocaleCanadianEnglish is the locale for Canadian English
	LocaleCanadianEnglish = "en-CA"

	//LocaleBritishEnglish is the locale for UK English
	LocaleBritishEnglish = "en-GB"

	//LocaleIndianEnglish is the locale for Indian English
	LocaleIndianEnglish = "en-IN"

	//LocaleAmericanEnglish is the locale for American English
	LocaleAmericanEnglish = "en-US"

	// LocaleJapanese is the locale for Japanese
	LocaleJapanese = "ja-JP"
)

// IsEnglish - Check if locale is set to english
func IsEnglish(locale string) bool {
	switch locale {
	case LocaleAmericanEnglish:
		return true
	case LocaleIndianEnglish:
		return true
	case LocaleBritishEnglish:
		return true
	case LocaleCanadianEnglish:
		return true
	case LocaleAustralianEnglish:
		return true
	default:
		return false
	}
}

// request

// Request is an Alexa skill request
// see https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#request-format
type Request struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Body    ReqBody `json:"request"`
	Context Context `json:"context"`
	Request struct {
		Type  string `json:"type"`
		Token string `json:"token"`
	}
}

// Session represents the Alexa skill session
type Session struct {
	New         bool   `json:"new"`
	SessionID   string `json:"sessionId"`
	Application struct {
		ApplicationID string `json:"applicationId"`
	} `json:"application"`
	Attributes map[string]interface{} `json:"attributes"`
	User       struct {
		UserID      string `json:"userId"`
		AccessToken string `json:"accessToken,omitempty"`
	} `json:"user"`
}

// Context represents the Alexa skill request context
type Context struct {
	System struct {
		APIAccessToken string `json:"apiAccessToken"`
		Device         struct {
			DeviceID string `json:"deviceId,omitempty"`
		} `json:"device,omitempty"`
		Application struct {
			ApplicationID string `json:"applicationId,omitempty"`
		} `json:"application,omitempty"`
		User struct {
			UserID      string `json:"userId"`
			AccessToken string `json:"accessToken,omitempty"`
		} `json:"user"`
	} `json:"System,omitempty"`
}

// ReqBody is the actual request information
type ReqBody struct {
	Type                 string `json:"type"`
	RequestID            string `json:"requestId"`
	Timestamp            string `json:"timestamp"`
	Locale               string `json:"locale"`
	Intent               Intent `json:"intent,omitempty"`
	Reason               string `json:"reason,omitempty"`
	DialogState          string `json:"dialogState,omitempty"`
	Token                string `json:"token"`
	OffsetInMilliseconds int64  `json:"offsetInMilliseconds"`
	Error                struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error"`
	Cause struct {
		RequestID string `json:"requestID"`
	} `json:"cause"`
}

// Intent is the Alexa skill intent
type Intent struct {
	Name  string          `json:"name"`
	Slots map[string]Slot `json:"slots"`
}

// Slot is an Alexa skill slot
type Slot struct {
	Name        string      `json:"name"`
	Value       string      `json:"value"`
	Resolutions Resolutions `json:"resolutions"`
}

// Resolutions - Alexa skill Resolutions
type Resolutions struct {
	ResolutionPerAuthority []struct {
		Values []struct {
			Value struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"value"`
		} `json:"values"`
	} `json:"resolutionsPerAuthority"`
}
