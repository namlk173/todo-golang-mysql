package entities

type Todo struct {
	id        int
	name      string
	expired   bool
	completed bool
}

// SETTER
func (t *Todo) SetId(id int) {
	t.id = id
}
func (t *Todo) SetName(name string) {
	t.name = name
}
func (t *Todo) SetExpired(isExpired bool) {
	t.expired = isExpired
}
func (t *Todo) SetCompleted(isCompleted bool) {
	t.completed = isCompleted
}

// GETTER
func (t Todo) GetId() int {
	return t.id
}
func (t Todo) GetName() string {
	return t.name
}
func (t Todo) GetExpired() bool {
	return t.expired
}
func (t Todo) GetCompleted() bool {
	return t.completed
}

// CONSTRUCTOR WITHIN ID
func NewTodo(id int, name string, expried, completed bool) *Todo {
	todo := new(Todo)
	todo.id = id
	todo.name = name
	todo.expired = expried
	todo.completed = completed
	return todo
}
