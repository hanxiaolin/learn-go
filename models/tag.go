package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by" xorm:"updated 'updated_at'"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func AddTag(name string, state int, createdBy string) (err error) {
	var tag = &Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	if err = db.Create(tag).Error; err != nil {
		return
	}

	return
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func DeleteTag(id int) bool {
	var tag Tag
	db.Where("id = ?", id).Delete(&tag)

	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}

func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

	return true
}

// BeforeCreate
//func (tag *Tag) BeforeCreate(scope *gorm.Scope) (err error) {
//	err = scope.SetColumn("CreatedOn", time.Now().Unix())
//	if err != nil {
//		return
//	}
//
//	err = scope.SetColumn("ModifiedOn", time.Now().Unix())
//	if err != nil {
//		return
//	}
//
//	return nil
//}
//
//func (tag *Tag) BeforeUpdate(scope *gorm.Scope) (err error) {
//	err = scope.SetColumn("ModifiedOn", time.Now().Unix())
//	if err != nil {
//		return
//	}
//
//	return nil
//}
