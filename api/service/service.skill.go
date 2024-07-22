package service

import (
	"fmt"
	"log"
	"net/http"

	errs "github.com/chayutK/skill-kafka-api/errors"
	"github.com/chayutK/skill-kafka-api/repository"
	"github.com/chayutK/skill-kafka-api/schemas"
	"github.com/gin-gonic/gin"
)

type ServiceSkill interface {
	GetByKey(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	UpdateByKey(ctx *gin.Context)
	UpdateName(ctx *gin.Context)
	UpdateDescription(ctx *gin.Context)
	UpdateLogo(ctx *gin.Context)
	UpdateTags(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type serviceSkill struct {
	repository repository.RepositorySkill
}

func NewServiceSkill(r repository.RepositorySkill) *serviceSkill {
	return &serviceSkill{repository: r}
}

func (ss *serviceSkill) GetByKey(ctx *gin.Context) {
	key := ctx.Param("key")
	fmt.Println(key)
	if key == "" {
		ctx.JSON(http.StatusBadRequest, errs.BadRequestError())
		return
	}
	skill, err := ss.repository.GetByKey(key)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errs.NotFoundError())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   skill,
	})
}

func (ss *serviceSkill) GetAll(ctx *gin.Context) {
	skills, err := ss.repository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   skills,
	})
}

func (ss *serviceSkill) Create(ctx *gin.Context) {
	skill := schemas.Skill{}
	err := ctx.ShouldBindJSON(&skill)

	if err != nil {
		log.Printf("Error while binding create request : %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError(err))
		return
	}

	err = ss.repository.Create(skill)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "request accepted",
	})
}

func (ss *serviceSkill) UpdateByKey(ctx *gin.Context) {
	key := ctx.Param("key")
	skill := schemas.Skill{}
	err := ctx.ShouldBindJSON(&skill)
	if err != nil {
		log.Printf("Error while binding update request : %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError(err))
		return
	}
	skill.Key = key

	err = ss.repository.UpdateByKey(skill)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "request accepted",
	})
}

func (ss *serviceSkill) UpdateName(ctx *gin.Context) {
	key := ctx.Param("key")
	skill := schemas.Skill{}
	err := ctx.ShouldBindJSON(&skill)
	if err != nil {
		log.Printf("Error while binding update name request : %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError(err))
		return
	}
	skill.Key = key

	err = ss.repository.UpdateName(skill)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "request accepted",
	})
}

func (ss *serviceSkill) UpdateTags(ctx *gin.Context) {
	key := ctx.Param("key")
	skill := schemas.Skill{}
	err := ctx.ShouldBindJSON(&skill)
	if err != nil {
		log.Printf("Error while binding update tags request : %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError(err))
		return
	}
	skill.Key = key

	err = ss.repository.UpdateTags(skill)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "request accepted",
	})
}

func (ss *serviceSkill) UpdateDescription(ctx *gin.Context) {
	key := ctx.Param("key")
	skill := schemas.Skill{}
	err := ctx.ShouldBindJSON(&skill)
	if err != nil {
		log.Printf("Error while binding update description request : %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError(err))
		return
	}
	skill.Key = key

	err = ss.repository.UpdateDescription(skill)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "request accepted",
	})
}

func (ss *serviceSkill) UpdateLogo(ctx *gin.Context) {
	key := ctx.Param("key")
	skill := schemas.Skill{}
	err := ctx.ShouldBindJSON(&skill)
	if err != nil {
		log.Printf("Error while binding update logo request : %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError(err))
		return
	}
	skill.Key = key

	err = ss.repository.UpdateLogo(skill)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "request accepted",
	})
}

func (ss *serviceSkill) Delete(ctx *gin.Context) {
	key := ctx.Param("key")

	err := ss.repository.Delete(key)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.InternalServerError(err))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "request accepted",
	})
}
