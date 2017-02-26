package greenergrass

var nameTestCases = []struct {
	input    string
	expected Name
}{
	{
		"John Moore",
		Name{first: "John", last: "Moore"},
	},
	{
		"Donald John Trump",
		Name{first: "Donald", middle: "John", last: "Trump"},
	},
	{
		"Hillary R. Clinton",
		Name{first: "Hillary", middle: "R.", last: "Clinton"},
	},
	{
		"",
		Name{},
	},
	{
		"Butherus",
		Name{first: "Butherus"},
	},
	{
		"Jose Gonzalez Quintero Perez",
		Name{first: "Jose", middle: "Gonzalez Quintero", last: "Perez"},
	},
	{
		"Walken, Christopher",
		Name{first: "Christopher", last: "Walken"},
	},
}
