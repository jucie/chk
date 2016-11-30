// Package chk contains some functions to verify the integrity of document ids.

package chk

import (
	"unicode"
)

func digit(s string, multiVector []int) (byte, bool) {
	sum := 0
	for i, mult := range multiVector {
		b := byte(s[i])
		if !unicode.IsDigit(rune(b)) {
			return byte('0'), false
		}
		sum += int(b-'0') * mult
	}
	dig := sum % 11
	if dig < 2 {
		return byte('0'), true
	}
	return byte((11 - dig) + '0'), true
}

func digitIsOk(s string, multiVector []int) bool {
	dig, ok := digit(s, multiVector)
	if !ok {
		return false
	}
	return dig == s[len(multiVector)]
}

var multCnpj = [...]int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

// CnpjIsValid returns true if the CNPJ is well formed and false otherwise.
// (CNPJ is brazilian company id).
func CnpjIsValid(cnpj string) bool {
	if len(cnpj) != 14 {
		return false
	}
	return digitIsOk(cnpj, multCnpj[1:]) && digitIsOk(cnpj, multCnpj[:])
}

var multCpf = [...]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

// CpfIsValid returns true if the CPF is well formed and false otherwise.
// (CPF is brazilian citizen id).
func CpfIsValid(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}
	return digitIsOk(cpf, multCpf[1:]) && digitIsOk(cpf, multCpf[:])
}

// Ean13IsValid returns true if the EAN-13 code is well formed and false otherwise.
// EAN-13 stands for European Article Number 13 digits.
// See http://pt.wikipedia.org/wiki/EAN-13
func Ean13IsValid(ean string) bool {
	if len(ean) != 13 {
		return false
	}
	b := []byte(ean)
	sum := make([]int, 2)
	for i := 0; i < 12; i++ {
		c := int(b[i])
		if !unicode.IsDigit(rune(c)) {
			return false
		}
		sum[i&1] += c - int('0')
	}
	total := sum[0] + sum[1]*3
	digit := (10 - (total % 10)) % 10
	return ean[12] == byte(digit+int('0'))
}
