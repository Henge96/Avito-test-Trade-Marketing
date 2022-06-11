package app

type App struct {
	repo Repo
}

func New(a Repo) *App {
	return &App{
		repo: a,
	}
}
