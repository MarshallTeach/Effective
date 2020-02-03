package content_service

import (
	"Effective/models"
	"Effective/models/content"
)

type Column struct {
	models.Model
	Name	string `json:"name"`
	ParentID	int `json:"parent_id"`
	CoverImageUrl string `json:"cover_image_url"`
}


func (c *Column) Get() ([]*content.Column, error) {
	var (
		columns []*content.Column
	)
	columns, err := content.GetColumn()
	if err != nil {
		return nil, err
	}

	return columns, nil
}