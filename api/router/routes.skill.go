package routes

import (
	"database/sql"

	"github.com/chayutK/skill-kafka-api/repository"
	"github.com/chayutK/skill-kafka-api/service"
	"github.com/gin-gonic/gin"
)

func InitSkillRouter(v *gin.RouterGroup, db *sql.DB) {
	repositorySkill := repository.NewRepositorySkill(db)
	serviceSkill := service.NewServiceSkill(repositorySkill)

	v.GET("/skills/:key", serviceSkill.GetByKey)
	v.GET("/skills", serviceSkill.GetAll)
	v.POST("/skills", serviceSkill.Create)
	v.PUT("/skills/:key", serviceSkill.UpdateByKey)
	v.PATCH("/skills/:key/actions/name", serviceSkill.UpdateName)
	v.PATCH("/skills/:key/actions/description", serviceSkill.UpdateDescription)
	v.PATCH("/skills/:key/actions/logo", serviceSkill.UpdateLogo)
	v.PATCH("/skills/:key/actions/tags", serviceSkill.UpdateTags)
	v.DELETE("/skills/:key", serviceSkill.Delete)
}
