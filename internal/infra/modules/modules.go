package modules

func InitializeModules(tickRate uint16) (ControllersModules, UseCasesModules, RepositoriesModules) {
	repositoriesModules := NewRepositoriesModules()
	useCasesModules := NewUseCaseModules(repositoriesModules, tickRate)
	controllersModules := NewControllersModules(useCasesModules)
	return controllersModules, useCasesModules, repositoriesModules
}
