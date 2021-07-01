package ihome

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIhomeCtomContentImgUploadAPIRequest
实拍图投稿链路上传图片 API请求
alibaba.ihome.ctom.content.img.upload

实拍图投稿链路上传图片 */
type AlibabaIhomeCtomContentImgUploadAPIRequest struct {
	model.Params
	// materialDTO
	_materialDto *UploadPicMaterialDto
}

// New
