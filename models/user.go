package models

var User *UserStatus

type UserStatus struct {
	ID           int
	Action       string
	CurrentPage  int
	selectedLine int
}

func (u *UserStatus) SetAction(action string) {
	u.Action = action
}

func (u *UserStatus) SetCurrentPage(page int) {
	u.CurrentPage = page
}

func (u *UserStatus) Clear() {
	u.Action = ""
	u.CurrentPage = 0
}

func (u *UserStatus) SelectLine(id int) {
	u.selectedLine = id
}

func (u *UserStatus) SelectedLine() int {
	return u.selectedLine
}

func (u *UserStatus) GetAction() string {
	return u.Action
}
