package datastore

import "github.com/ogra1/fabrica/domain"

// Datastore interface for the database logic
type Datastore interface {
	BuildList() ([]domain.Build, error)
	BuildCreate(name, repo, branch string) (string, error)
	BuildUpdate(id, status string, duration int) error
	BuildUpdateDownload(id, download string) error
	BuildUpdateContainer(id, container string) error
	BuildGet(id string) (domain.Build, error)
	BuildDelete(id string) error
	BuildListForRepo(name, branch string) ([]domain.Build, error)

	BuildLogCreate(id, message string) error
	BuildLogList(id string) ([]domain.BuildLog, error)

	RepoCreate(name, repo, branch, keyID string) (string, error)
	RepoGet(id string) (domain.Repo, error)
	RepoList(watch bool) ([]domain.Repo, error)
	RepoUpdateHash(id, hash string) error
	RepoDelete(id string) error

	KeysCreate(name, data, password string) (string, error)
	KeysGet(id string) (domain.Key, error)
	KeysList() ([]domain.Key, error)
	KeysDelete(id string) error

	SettingsCreate(key, name, data string) (string, error)
	SettingsGet(key, name string) (domain.ConfigSetting, error)
}
