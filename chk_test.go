package chk

import (
	"testing"
)

func TestCnpjIsValid(t *testing.T) {
	if !CnpjIsValid("52803004000101") {
		t.Errorf("52803004000101 is a valid CNPJ but CnpjIsValid returned false")
	}
	if !CnpjIsValid("47508411105179") {
		t.Errorf("47508411105179 is a valid CNPJ but CnpjIsValid returned false")
	}
	if CnpjIsValid("123435678901234") {
		t.Errorf("123435678901234 is not a valid CNPJ but CnpjIsValid returned true")
	}
}

func TestCpfIsValid(t *testing.T) {
	if !CpfIsValid("84769380399") {
		t.Errorf("84769380399 is a valid CPF but CpfIsValid returned false")
	}
	if !CpfIsValid("43881322175") {
		t.Errorf("43881322175 is a valid CPF but CpfIsValid returned false")
	}
	if CpfIsValid("07203833400") {
		t.Errorf("07203833400 is not a valid CPF but CpfIsValid returned true")
	}
}
