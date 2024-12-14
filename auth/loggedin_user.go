package auth

type LoggedInUser struct {
}

func NewLoggedInUser() *LoggedInUser {
	return &LoggedInUser{}
}

func (l *LoggedInUser) UserId() int {
	return 1
}
