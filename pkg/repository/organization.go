package repository

import (
	"Tryscon-api/models"
	"github.com/jmoiron/sqlx"
)

type OrganizationRepository struct {
	db *sqlx.DB
}

func (o *OrganizationRepository) CreateOrganization(db *sqlx.DB, input models.CreateOrganization) (models.CreateOrganization, error) {
	var organization models.CreateOrganization

	err := db.Exec()

	return organization, nil
}
