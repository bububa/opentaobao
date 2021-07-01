package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuPhotobankUploadAPIRequest
图片银行图片上传开放接口 API请求
alibaba.icbu.photobank.upload

图片银行图片上传开放接口 */
type AlibabaIcbuPhotobankUploadAPIRequest struct {
	model.Params
	// 上传图片名称
	_fileName string
	// 上传图片所在分组
	_groupId string
	// 图片字节数组
	_imageBytes *model.File
	// 扩展参数信息,如ICVID
	_extraContext string
}

// New
