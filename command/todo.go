package command

import "time"

// Format represents how timestamps show look
var Format = "2006-01-02T15:04"

// Todoer implements the functionality to view todos
type Todoer interface {
	byName()
	byFile()
	byTimestamp()
	byWeight()
}

// Todo represents a todo
type Todo struct {
	User      string
	File      string
	Message   string
	Timestamp time.Time
	Weight    int
}

// NewTodo creates a new reference of a todo
func NewTodo(user, msg, ts, file string, weight int) (*Todo, error) {
	t := &Todo{
		User:    user,
		File:    file,
		Message: msg,
		Weight:  weight,
	}
	timestamp, err := time.Parse(Format, ts)
	if err != nil {
		return nil, err
	}
	t.Timestamp = timestamp
	return t, err
}

// UserTodos is a slice type made for easier sorting
type UserTodos []Todo

// Len gets the length of the slice
func (u UserTodos) Len() int {
	return len(u)
}

// Less does a comparison of the 2 given arguments
func (u UserTodos) Less(i, j int) bool {
	return u[i].User < u[j].User
}

// Swap switchs the place in the slice for the 2 given arguments
func (u UserTodos) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// FileTodos is a slice type made for easier sorting
type FileTodos []Todo

// Len gets the length of the slice
func (f FileTodos) Len() int {
	return len(f)
}

// Less does a comparison of the 2 given arguments
func (f FileTodos) Less(i, j int) bool {
	return f[i].File < f[j].File
}

// Swap switchs the place in the slice for the 2 given arguments
func (f FileTodos) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// TimestampTodos is a slice type made for easier sorting
type TimestampTodos []Todo

// Len gets the length of the slice
func (t TimestampTodos) Len() int {
	return len(t)
}

// Less does a comparison of the 2 given arguments
func (t TimestampTodos) Less(i, j int) bool {
	return t[i].User < t[j].User
}

// Swap switchs the place in the slice for the 2 given arguments
func (t TimestampTodos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// WeightTodos is a slice type made for easier sorting
type WeightTodos []Todo

// Len gets the length of the slice
func (w WeightTodos) Len() int {
	return len(w)
}

// Less does a comparison of the 2 given arguments
func (w WeightTodos) Less(i, j int) bool {
	return w[i].User > w[j].User
}

// Swap switchs the place in the slice for the 2 given arguments
func (w WeightTodos) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
