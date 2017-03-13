package greenergrass

var nameTestCases = []struct {
	input    string
	expected Name
}{
	{
		"John Moore",
		Name{First: "John", Last: "Moore"},
	},
	{
		"Donald John Trump",
		Name{First: "Donald", Middle: "John", Last: "Trump"},
	},
	{
		"Hillary R. Clinton",
		Name{First: "Hillary", Middle: "R.", Last: "Clinton"},
	},
	{
		"",
		Name{},
	},
	{
		"Butherus",
		Name{First: "Butherus"},
	},
	{
		"Jose Gonzalez Quintero Perez",
		Name{First: "Jose", Middle: "Gonzalez Quintero", Last: "Perez"},
	},
	{
		"Walken, Christopher",
		Name{First: "Christopher", Last: "Walken"},
	},
	{
		"Agent Frank Horrigan",
		Name{First: "Frank", Last: "Horrigan", Prefix: "Agent"},
	},
	{
		"Sir Paul McCartney",
		Name{First: "Paul", Last: "McCartney", Prefix: "Sir"},
	},
	{
		"Robert Downey Jr.",
		Name{First: "Robert", Last: "Downey", Suffix: "Jr."},
	},
	{
		"King George III",
		Name{First: "George", Prefix: "King", Suffix: "III"},
	},
}
