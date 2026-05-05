package enums

type Country string

const (
	Guatemala   Country = "GT"
	Mexico      Country = "MX"
	CostaRica   Country = "CR"
	ElSalvador  Country = "SV"
	Honduras    Country = "HN"
	Nicaragua   Country = "NI"
	Panama      Country = "PA"
	Colombia    Country = "CO"
	Peru        Country = "PE"
	Chile       Country = "CL"
	Argentina   Country = "AR"
	Uruguay     Country = "UY"
	Paraguay    Country = "PY"
	Bolivia     Country = "BO"
	Ecuador     Country = "EC"
	Venezuela   Country = "VE"
	Brazil      Country = "BR"
	DominicanRp Country = "DO"
	Spain       Country = "ES"
	USA         Country = "US"
)

// String returns the string representation of the Country.
func (c Country) String() string {
	return string(c)
}
