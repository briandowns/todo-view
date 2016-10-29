package command

// Priorities represents the group of priorities TODO-View
// represents
var Priorities = map[int]string{
	5: "Not at all important",
	4: "Slightly important",
	3: "Moderately important",
	2: "Somewhat important",
	1: "Extremely important",
}

// Priority holds an individual priority
type Priority struct {
	Value       int
	Description string
}
