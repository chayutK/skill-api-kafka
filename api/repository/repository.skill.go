package repository

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/chayutK/skill-kafka-api/producer"
	"github.com/chayutK/skill-kafka-api/schemas"
	"github.com/lib/pq"
)

type RepositorySkill interface {
	GetByKey(key string) (schemas.Skill, error)
	GetAll() ([]schemas.Skill, error)
	Create(skill schemas.Skill) error
	UpdateByKey(skill schemas.Skill) error
	UpdateName(skill schemas.Skill) error
	UpdateDescription(skill schemas.Skill) error
	UpdateLogo(skill schemas.Skill) error
	UpdateTags(skill schemas.Skill) error
	Delete(key string) error
}

type repositorySkill struct {
	db *sql.DB
}

func NewRepositorySkill(db *sql.DB) *repositorySkill {
	return &repositorySkill{db: db}
}

func (rs *repositorySkill) GetByKey(key string) (schemas.Skill, error) {
	result := schemas.Skill{}
	q := "SELECT key,name,description,logo,tags FROM skill WHERE key=$1"
	row := rs.db.QueryRow(q, key)

	err := row.Scan(&result.Key, &result.Name, &result.Description, &result.Logo, pq.Array(&result.Tags))
	if err != nil {
		log.Printf("Error while get skill by key : %s", err.Error())
		return result, err
	}

	return result, nil
}

func (rs *repositorySkill) GetAll() ([]schemas.Skill, error) {
	result := []schemas.Skill{}

	q := "SELECT * FROM skill"
	rows, err := rs.db.Query(q)
	if err != nil {
		log.Printf("Error while get skill by key : %s", err.Error())
		return result, err
	}

	for rows.Next() {
		skill := schemas.Skill{}
		err := rows.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
		if err != nil {
			log.Printf("Error while get skill by key : %s", err.Error())
			return result, err
		}
		result = append(result, skill)
		// fmt.Println(result, rows.Next())
	}

	return result, nil
}

func (rs *repositorySkill) Create(skill schemas.Skill) error {
	message, err := json.Marshal(skill)
	if err != nil {
		log.Printf("Error while marshal skill creating : %s", err.Error())
		return err
	}
	producer.SendMessage("Create", message)

	return nil
}

func (rs *repositorySkill) UpdateByKey(skill schemas.Skill) error {
	message, err := json.Marshal(skill)
	if err != nil {
		log.Printf("Error while marshal skill creating : %s", err.Error())
		return err
	}
	producer.SendMessage("Update", message)

	return nil
}

func (rs *repositorySkill) UpdateName(skill schemas.Skill) error {
	message, err := json.Marshal(skill)
	if err != nil {
		log.Printf("Error while marshal skill creating : %s", err.Error())
		return err
	}
	producer.SendMessage("UpdateName", message)

	return nil
}

func (rs *repositorySkill) UpdateDescription(skill schemas.Skill) error {
	message, err := json.Marshal(skill)
	if err != nil {
		log.Printf("Error while marshal skill creating : %s", err.Error())
		return err
	}
	producer.SendMessage("UpdateDesc", message)

	return nil
}

func (rs *repositorySkill) UpdateLogo(skill schemas.Skill) error {
	message, err := json.Marshal(skill)
	if err != nil {
		log.Printf("Error while marshal skill creating : %s", err.Error())
		return err
	}
	producer.SendMessage("UpdateLogo", message)

	return nil
}

func (rs *repositorySkill) UpdateTags(skill schemas.Skill) error {
	message, err := json.Marshal(skill)
	if err != nil {
		log.Printf("Error while marshal skill creating : %s", err.Error())
		return err
	}
	producer.SendMessage("UpdateTags", message)

	return nil
}

func (rs *repositorySkill) Delete(key string) error {
	message, err := json.Marshal(key)
	if err != nil {
		log.Printf("Error while marshal skill creating : %s", err.Error())
		return err
	}
	producer.SendMessage("Delete", message)

	return nil
}
