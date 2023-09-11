package application

import "volga-core/types"

func CreateApplication(app Application, session types.Session) (Application, error) {
	app.User = session.Username
	return Create(app)
}

func UpdateApplication(code string, app Application, session types.Session) (Application, error) {
	app.User = session.Username
	return UpdateByCode(code, app)
}

func ListUserApplication(session types.Session) ([]Application, error) {
	return ListByUser(session.Username)
}
