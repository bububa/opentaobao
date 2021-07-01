package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterPictureUploadAPIRequest
上传图片 API请求
tmall.servicecenter.picture.upload

给服务商ERP系统使用，用于上传图片保存在天猫，一般用于工单信息回传时候保存服务商的服务证明信息相关的图片。 */
type TmallServicecenterPictureUploadAPIRequest struct {
	model.Params
	// 图片文件二进制流
	_img *model.File
	// 图片全称包括扩展名。目前支持 jpg jpeg png
	_pictureName string
	// true返回Https地址
	_isHttps bool
}

// New
