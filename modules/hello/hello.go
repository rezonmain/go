package hello

import (
	"rsc.io/quote/v3"
)

func Hello() string {
	return hello()
}

func Proverb() string {
	return quote.Concurrency()
}

func hello() string {
	return quote.HelloV3()
}
