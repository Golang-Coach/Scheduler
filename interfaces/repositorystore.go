package interfaces

import (
	"github.com/Golang-Coach/Scheduler/models"
)

type IRepositoryStore interface {
	AddPackage(repositoryInfo models.RepositoryInfo) error
	UpdatePackage(pack *models.RepositoryInfo) error
	FindPackage(query interface{}) (*models.RepositoryInfo, error)
	FindPackageWithinLimit(query interface{}, skip int, limit int) (*[]models.RepositoryInfo, error)
}
