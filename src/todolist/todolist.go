/*
TodoList help to do tascks of man
*/
package todolist


type user struct {
	name string
}

func NewUser(n string) *user{
	return &user{name: n}
}

//Getter user name
func (u *user) GetUser() string {
	return u.name
}

//Setter user name
func (u *user) Name(n string) {
	u.name = n
}

type Task struct {
	title        string
	description string
	created string
}

//getter task title
func (t *Task) Title() string{
return t.title
}

//getter task description
func (t *Task) Description() string{
	return t.description
	}

//getter task created
func (t *Task) Created() string{
	return t.created
	}

//setter task title
func (t *Task) SetTitle(title string){
	t.title = title
	}

//setter task description
func (t *Task) SetDescription(d string){
	t.description = d
	}

//setter task created
func (t *Task) SetCreated(c string){
	t.created = c
	}