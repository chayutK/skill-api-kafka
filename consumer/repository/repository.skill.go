package repository

import (
	"database/sql"

	"github.com/chayutK/skill-api-kafka/comsumer/schemas"
	"github.com/lib/pq"
)

type RepositorySkill interface {
	Create(skill schemas.Skill) error
	UpdateByKey(skill schemas.Skill) error
	UpdateName(skill schemas.Skill) error
	UpdateDescription(skill schemas.Skill) error
	UpdateLogo(skill schemas.Skill) error
	UpdateTags(skill schemas.Skill) error
	// UpdateDescription(key, description string) (*schemas.Skill, error)
	// UpdateLogo(key, logo string) (*schemas.Skill, error)
	// UpdateTags(key string, tags []string) (*schemas.Skill, error)
	Delete(key string) error
}

type repositorySkill struct {
	db *sql.DB
}

func NewRepositorySkill(db *sql.DB) *repositorySkill {
	return &repositorySkill{db: db}
}

func (rs *repositorySkill) Create(skill schemas.Skill) error {
	q := "INSERT INTO skill (key,name,description,logo,tags) values ($1,$2,$3,$4,$5) RETURNING key,name,description,logo,tags"
	row := rs.db.QueryRow(q, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))

	result := schemas.Skill{}
	err := row.Scan(&result.Key, &result.Name, &result.Description, &result.Logo, pq.Array(&result.Tags))
	if err != nil {
		return err
	}

	return nil
}

func (rs *repositorySkill) UpdateByKey(skill schemas.Skill) error {
	q := "UPDATE skill SET name = $2,description = $3, logo=$4,tags=$5 WHERE key=$1 "
	_, err := rs.db.Exec(q, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))
	if err != nil {
		return err
	}

	return nil
}

func (rs *repositorySkill) UpdateName(skill schemas.Skill) error {
	q := "UPDATE skill SET name = $2 WHERE key=$1 "
	_, err := rs.db.Exec(q, skill.Key, skill.Name)
	if err != nil {
		return err
	}
	return nil
}

func (rs *repositorySkill) UpdateDescription(skill schemas.Skill) error {
	q := "UPDATE skill SET description = $2 WHERE key=$1 "
	_, err := rs.db.Exec(q, skill.Key, skill.Description)
	if err != nil {
		return err
	}
	return nil
}

func (rs *repositorySkill) UpdateLogo(skill schemas.Skill) error {
	q := "UPDATE skill SET logo = $2 WHERE key=$1 "
	_, err := rs.db.Exec(q, skill.Key, skill.Logo)
	if err != nil {
		return err
	}
	return nil
}

func (rs *repositorySkill) UpdateTags(skill schemas.Skill) error {
	q := "UPDATE skill SET tags = $2 WHERE key=$1 "
	_, err := rs.db.Exec(q, skill.Key, pq.Array(skill.Tags))
	if err != nil {
		return err
	}
	return nil
}

func (rs *repositorySkill) Delete(key string) error {
	q := "DELETE FROM skill WHERE key=$1"
	_, err := rs.db.Exec(q, key)
	if err != nil {
		return err
	}

	return nil
}
