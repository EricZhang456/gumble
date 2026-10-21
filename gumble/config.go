package gumble

import (
	"time"
)

type ClientVersion interface {
	ToVersionInt() any
	Major() uint16
	Minor() uint16
	Patch() uint16
}

type ClientVersionLegacy struct {
	major uint16
	minor uint8
	patch uint8
}

type ClientVersionNew struct {
	major uint16
	minor uint16
	patch uint16
	build uint16
}

// Config holds the Mumble configuration used by Client. A single Config should
// not be shared between multiple Client instances.
type Config struct {
	// User name used when authenticating with the server.
	Username string
	// Password used when authenticating with the server. A password is not
	// usually required to connect to a server.
	Password string
	// The initial access tokens to the send to the server. Access tokens can be
	// resent to the server using:
	//  client.Send(config.Tokens)
	Tokens AccessTokens

	// AudioInterval is the interval at which audio packets are sent. Valid
	// values are: 10ms, 20ms, 40ms, and 60ms.
	AudioInterval time.Duration
	// AudioDataBytes is the number of bytes that an audio frame can use.
	AudioDataBytes int

	// The event listeners used when client events are triggered.
	Listeners      Listeners
	AudioListeners AudioListeners

	Release string

	Version ClientVersion
}

// NewConfig returns a new Config struct with default values set.
func NewConfig() *Config {
	return &Config{
		AudioInterval:  AudioDefaultInterval,
		AudioDataBytes: AudioDefaultDataBytes,
		Release:        "gumble",
		Version: &ClientVersionNew{
			major: 1,
			minor: 5,
			patch: 0,
			build: 0,
		},
	}
}

// Attach is an alias of c.Listeners.Attach.
func (c *Config) Attach(l EventListener) Detacher {
	return c.Listeners.Attach(l)
}

// AttachAudio is an alias of c.AudioListeners.Attach.
func (c *Config) AttachAudio(l AudioListener) Detacher {
	return c.AudioListeners.Attach(l)
}

// AudioFrameSize returns the appropriate audio frame size, based off of the
// audio interval.
func (c *Config) AudioFrameSize() int {
	return int(c.AudioInterval/AudioDefaultInterval) * AudioDefaultFrameSize
}

func (v ClientVersionLegacy) ToVersionInt() any {
	var value uint32
	value = (uint32(v.major) << 16) | (uint32(v.minor) << 8) | uint32(v.patch)&0xFF
	return value
}

func (v ClientVersionLegacy) Major() uint16 {
	return v.major
}

func (v ClientVersionLegacy) Minor() uint16 {
	return uint16(v.minor)
}

func (v ClientVersionLegacy) Patch() uint16 {
	return uint16(v.patch)
}

func (v ClientVersionNew) ToVersionInt() any {
	var value uint64
	value = ((uint64(v.major) & 0xFFFF) << 48) | ((uint64(v.minor) & 0xFFFF) << 32) |
		((uint64(v.minor) & 0xFFFF) << 16) | uint64(v.patch)&0xFFFF
	return value
}

func (v ClientVersionNew) Major() uint16 {
	return v.major
}

func (v ClientVersionNew) Minor() uint16 {
	return v.minor
}

func (v ClientVersionNew) Patch() uint16 {
	return v.patch
}
