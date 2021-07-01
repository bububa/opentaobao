package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaGspSupplyImageUploadAPIRequest
gsp图片上传 API请求
alibaba.gsp.supply.image.upload

上传图片至目标海外平台的素材空间 */
type AlibabaGspSupplyImageUploadAPIRequest struct {
	model.Params
	// 图片名称
	_fileName string
	// 图片文件流，像素宽度不小于500，不大于2000，像素长度不小于500，不大于2000
	_fileContent *model.File
}

// New
