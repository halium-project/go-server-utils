package is

import (
	"errors"
	"regexp"

	"github.com/asaskevich/govalidator"
)

type stringValidator func(string) bool

func newStringRule(validator stringValidator) func(string) error {
	return func(str string) error {
		if validator(str) {
			return nil
		}

		return errors.New(InvalidFormat)
	}
}

var (
	DNSName          = newStringRule(govalidator.IsDNSName)
	Email            = newStringRule(govalidator.IsEmail)
	Host             = newStringRule(govalidator.IsHost)
	ID               = newStringRule(govalidator.IsUUIDv4)
	IPv4             = newStringRule(govalidator.IsIPv4)
	IPv6             = newStringRule(govalidator.IsIPv6)
	LowerCase        = newStringRule(govalidator.IsLowerCase)
	Port             = newStringRule(govalidator.IsPort)
	UpperCase        = newStringRule(govalidator.IsUpperCase)
	URL              = newStringRule(govalidator.IsURL)
	UTFDigit         = newStringRule(govalidator.IsUTFDigit)
	UTFLetter        = newStringRule(govalidator.IsUTFLetter)
	UTFLetterNumeric = newStringRule(govalidator.IsUTFLetterNumeric)
)

func StringInRange(min int, max int) func(string) error {
	return func(str string) error {
		if len(str) < min {
			return errors.New(TooShort)
		}

		if len(str) > max {
			return errors.New(TooLong)
		}

		return nil
	}
}

func OnOfString(matches ...string) func(string) error {
	return func(str string) error {
		for _, match := range matches {
			if str == match {
				return nil
			}
		}

		return errors.New(UnexpectedValue)
	}
}

func MatchingString(pattern string) func(string) error {
	return func(str string) error {
		match, err := regexp.MatchString(pattern, str)
		if err != nil {
			panic(err)
		}

		if match {
			return nil
		}

		return errors.New(InvalidFormat)
	}
}
