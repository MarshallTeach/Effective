package content

import (
	"Effective/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Column struct {
	models.Model
	Name	string `json:"name"`
	ParentID	int `json:"parent_id"`
	CoverImageUrl string `json:"cover_image_url"`
}



func AddColumn(data map[string]interface{}) error {
	column := Column{
		Name:          data["name"].(string),
		ParentID:      data["parent_id"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}
	
	if err := db.Create(&column).Error; err != nil {
		return err
	}
	
	return nil
}

func DeleteColumn(ids []int) error {
	for _, id := range ids {
		if err := db.Where("id = ?", id).Delete(Column{}).Error; err != nil {
			return err
		}
	}
	
	return nil
}

func EditColumn(id int, data interface{}) error {
	if err := db.Model(&Column{}).Where("id = ? AND deleted_at = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}
	
	return nil
}

func GetColumn() ([]*Column, error) {
	var column []*Column
	if err := db.Find(&Column{}).Error; err != nil{
		return nil, err
	}
	return column, nil
}