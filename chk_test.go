package chk

import (
	"testing"
)

func TestCnpjIsValid(t *testing.T) {
	valid := []string{
		"01569444000105",
		"04776111000145",
		"04949587000130",
		"07319327000170",
		"07846559000186",
		"10381592000192",
		"14661167000108",
		"21100656000156",
		"22702850000174",
		"51015212000183",
		"60570793000106",
		"65370025000122",
		"67756288000108",
		"86526696000100",
		"52803004000101",
		"47508411105179",
	}

	for _, s := range valid {
		if !CnpjIsValid(s) {
			t.Errorf(s + " is a valid CNPJ but CnpjIsValid returned false")
		}
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

func TestEan13IsValid(t *testing.T) {
	if !Ean13IsValid("7899999912349") {
		t.Errorf("7899999912349 is a valid EAN-13 but Ean13IsValid returned false")
	}
	if !Ean13IsValid("7891627314051") {
		t.Errorf("7891627314051 is a valid EAN-13 but Ean13IsValid returned false")
	}
	if Ean13IsValid("4388132217238") {
		t.Errorf("4388132217238 is not a valid EAN-13 but Ean13IsValid returned true")
	}
}
