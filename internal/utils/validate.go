package utils

import (
	"errors"
	"regexp"

	"github.com/m16yusuf/belajar-gin/internal/models"
)

func ValidateBody(ping models.Ping) error {
	if ping.Id <= 0 {
		return errors.New("id tidak boleh dibawah 0")
	}
	if len(ping.Message) < 8 {
		return errors.New("panjang pesan harus diatas 8 karakter")
	}
	re, err := regexp.Compile("^[lLpPmMfF]$")
	if err != nil {
		return err
	}
	if isMatched := re.Match([]byte(ping.Gender)); !isMatched {
		return errors.New("gender harus berisikan huruf l, L, m, M, f, F, p, P")
	}
	return nil
}
