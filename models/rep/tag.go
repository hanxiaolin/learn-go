package rep

type AddTagReq struct {
	Name      string `form:"name" binding:"required"`
	State     int    `form:"state" binding:"required"`
	CreatedBy string `form:"created_by" binding:"required"`
}

type EditTagReq struct {
	Id         int    `form:"id" binding:"required"`
	Name       string `form:"name" binding:"required"`
	State      int    `form:"state" binding:"required"`
	NodifiedBy string `form:"modified_by" binding:"required"`
}
