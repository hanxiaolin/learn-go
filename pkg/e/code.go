package e

type ResponseType int

const (
	SUCCESS        ResponseType = 200
	ERROR          ResponseType = 500
	INVALID_PARAMS ResponseType = 400

	ERROR_EXIST_TAG         ResponseType = 10001
	ERROR_NOT_EXIST_TAG     ResponseType = 10002
	ERROR_NOT_EXIST_ARTICLE ResponseType = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    ResponseType = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT ResponseType = 20002
	ERROR_AUTH_TOKEN               ResponseType = 20003
	ERROR_AUTH                     ResponseType = 20004

	// 保存图片失败
	ERROR_UPLOAD_SAVE_IMAGE_FAIL ResponseType = 30001
	// 检查图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FAIL ResponseType = 30002
	// 校验图片错误，图片格式或大小有问题
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT ResponseType = 30003

	ERROR_CHECK_EXIST_ARTICLE_FAIL ResponseType = 30004
	ERROR_GET_ARTICLE_FAIL         ResponseType = 30005
)
