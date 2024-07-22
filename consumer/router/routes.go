package router

import (
	"database/sql"

	"github.com/chayutK/skill-api-kafka/comsumer/repository"
	"github.com/chayutK/skill-api-kafka/comsumer/service"
)

type Router interface {
	Route(key string, message []byte)
}

type router struct {
	db *sql.DB
}

func NewRouter(db *sql.DB) *router {
	return &router{
		db: db,
	}
}

func (r *router) Route(key string, message []byte) {
	repositorySkill := repository.NewRepositorySkill(r.db)
	serviceSkill := service.NewServiceSkill(repositorySkill)

	switch key {
	case "Create":
		serviceSkill.Create(message)
	case "Update":
		serviceSkill.UpdateByKey(message)
	case "UpdateName":
		serviceSkill.UpdateName(message)
	case "UpdateDesc":
		serviceSkill.UpdateDescription(message)
	case "UpdateLogo":
		serviceSkill.UpdateLogo(message)
	case "UpdateTags":
		serviceSkill.UpdateTags(message)
	case "Delete":
		serviceSkill.Delete(message)
	default:
	}

}
