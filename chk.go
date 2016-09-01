// Contains some functions to verify the integrity of document ids.

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

// returns true if the CNPJ is well formed and false otherwise
// (CNPJ is brazilian company id).
func CnpjIsValid(cnpj string) bool {
	if len(cnpj) != 14 {
		return false
	}
	return digitIsOk(cnpj, multCnpj[1:]) && digitIsOk(cnpj, multCnpj[:])
}

var multCpf = [...]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

// returns true if the CPF is well formed and false otherwise
// (CPF is brazilian citizen id).
func CpfIsValid(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}
	return digitIsOk(cpf, multCpf[1:]) && digitIsOk(cpf, multCpf[:])
}
