package enums

type Currency string

const (
	GTQ Currency = "GTQ"
	MXN Currency = "MXN"
	CRC Currency = "CRC"
	SVC Currency = "SVC"
	HNL Currency = "HNL"
	NIO Currency = "NIO"
	PAB Currency = "PAB"
	COP Currency = "COP"
	PEN Currency = "PEN"
	CLP Currency = "CLP"
	ARS Currency = "ARS"
	UYU Currency = "UYU"
	PYG Currency = "PYG"
	BOB Currency = "BOB"
	VEF Currency = "VEF" // or VES
	BRL Currency = "BRL"
	DOP Currency = "DOP"
	EUR Currency = "EUR"
	USD Currency = "USD"
)

// String returns the string representation of the Currency.
func (c Currency) String() string {
	return string(c)
}
