package config

import (
	u "exchangerate/utils"
)

var Spider *spider

type spider struct {
	BankIndonesiaUri string
	OrtaxUri         string
}

func init() {
	Spider = &spider{
		BankIndonesiaUri: u.Env("SPIDER_BANK_INDONESIA_URI"),
		OrtaxUri:         u.Env("SPIDER_ORTAX_URI"),
	}
}
