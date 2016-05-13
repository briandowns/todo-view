package command

var Weights = map[int]string{
	5: "Not at all important",
	4: "Slightly important",
	3: "Moderately important",
	2: "Somewhat important",
	1: "Extremely important",
}

// Weight is holds a weight
type Weight struct {
	Value       int
	Description string
}
