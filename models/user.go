package models

var User *UserStatus

type UserStatus struct {
	ID          int
	Action      string
	CurrentPage int
}

func (u UserStatus) SetAction(action string) {
	u.Action = action
}

func (u UserStatus) SetCurrentPage(page int) {
	u.CurrentPage = page
}

func (u UserStatus) Clear() {
	u.Action = ""
	u.CurrentPage = 0
}
