package v1

import (
	"code.shihuo.cn/gin-demo/logging"
	"code.shihuo.cn/gin-demo/models"
	"code.shihuo.cn/gin-demo/models/rep"
	"code.shihuo.cn/gin-demo/pkg/e"
	"code.shihuo.cn/gin-demo/pkg/setting"
	"code.shihuo.cn/gin-demo/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	var (
		err error
	)

	req := new(rep.AddTagReq)

	code := e.INVALID_PARAMS
	if err = c.ShouldBind(req); err == nil {
		if !models.ExistTagByName(req.Name) {
			code = e.SUCCESS
			err = models.AddTag(req.Name, req.State, req.CreatedBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	if err != nil {
		logging.Error(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

// @Summary 修改文章标签
// @Produce  json
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [put]
func EditTag(c *gin.Context) {
	var (
		err error
	)

	req := new(rep.EditTagReq)

	code := e.INVALID_PARAMS
	if err = c.Bind(req); err == nil {
		code = e.SUCCESS
		if models.ExistTagByID(req.Id) {
			data := make(map[string]interface{})
			data["modified_by"] = req.NodifiedBy
			if req.Name != "" {
				data["name"] = req.Name
			}
			if req.State != -1 {
				data["state"] = req.State
			}

			models.EditTag(req.Id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
