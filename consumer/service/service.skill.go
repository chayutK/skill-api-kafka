package service

import (
	"encoding/json"
	"log"

	"github.com/chayutK/skill-api-kafka/comsumer/repository"
	"github.com/chayutK/skill-api-kafka/comsumer/schemas"
)

type ServiceSkill interface {
	Create(message []byte)
	UpdateByKey(message []byte)
	UpdateName(message []byte)
	UpdateDescription(message []byte)
	UpdateLogo(message []byte)
	UpdateTags(message []byte)
	Delete(message []byte)
}

type serviceSkill struct {
	repository repository.RepositorySkill
}

func NewServiceSkill(r repository.RepositorySkill) *serviceSkill {
	return &serviceSkill{repository: r}
}

func (ss *serviceSkill) Create(message []byte) {
	skill := schemas.Skill{}
	err := json.Unmarshal(message, &skill)

	if err != nil {
		log.Printf("Error while binding create request : %s", err.Error())
		return
	}

	err = ss.repository.Create(skill)

	if err != nil {
		return
	}
}

func (ss *serviceSkill) UpdateByKey(message []byte) {
	skill := schemas.Skill{}
	err := json.Unmarshal(message, &skill)

	if err != nil {
		log.Printf("Error while binding update request : %s", err.Error())
	}

	err = ss.repository.UpdateByKey(skill)

	if err != nil {
		log.Printf("Error while updating skill: %s", err.Error())
	}
}

func (ss *serviceSkill) UpdateName(message []byte) {
	skill := schemas.Skill{}
	err := json.Unmarshal(message, &skill)

	if err != nil {
		log.Printf("Error while binding update name request : %s", err.Error())
	}

	err = ss.repository.UpdateName(skill)

	if err != nil {
		log.Printf("Error while updating skill's name : %s", err.Error())
	}
}

func (ss *serviceSkill) UpdateDescription(message []byte) {
	skill := schemas.Skill{}
	err := json.Unmarshal(message, &skill)

	if err != nil {
		log.Printf("Error while binding update description request : %s", err.Error())
	}

	err = ss.repository.UpdateDescription(skill)

	if err != nil {
		log.Printf("Error while updating skill's description : %s", err.Error())
	}
}

func (ss *serviceSkill) UpdateLogo(message []byte) {
	skill := schemas.Skill{}
	err := json.Unmarshal(message, &skill)

	if err != nil {
		log.Printf("Error while binding update logo request : %s", err.Error())
	}

	err = ss.repository.UpdateLogo(skill)

	if err != nil {
		log.Printf("Error while updating skill's logo : %s", err.Error())
	}
}

func (ss *serviceSkill) UpdateTags(message []byte) {
	skill := schemas.Skill{}
	err := json.Unmarshal(message, &skill)

	if err != nil {
		log.Printf("Error while binding update tags request : %s", err.Error())
	}

	err = ss.repository.UpdateTags(skill)

	if err != nil {
		log.Printf("Error while updating skill's tags : %s", err.Error())
	}
}

func (ss *serviceSkill) Delete(message []byte) {
	var key string
	err := json.Unmarshal(message, &key)
	if err != nil {
		log.Printf("Error while binding delete request : %s", err.Error())
	}

	err = ss.repository.Delete(key)

	if err != nil {
		log.Printf("Error while delete skill : %s", err.Error())
	}
}
