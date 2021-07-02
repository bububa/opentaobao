package ihome

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIhomeCtomContentImgUploadAPIRequest 实拍图投稿链路上传图片 API请求
// alibaba.ihome.ctom.content.img.upload
//
// 实拍图投稿链路上传图片
type AlibabaIhomeCtomContentImgUploadAPIRequest struct {
	model.Params
	// materialDTO
	_materialDto *UploadPicMaterialDto
}

// NewAlibabaIhomeCtomContentImgUploadRequest 初始化AlibabaIhomeCtomContentImgUploadAPIRequest对象
func NewAlibabaIhomeCtomContentImgUploadRequest() *AlibabaIhomeCtomContentImgUploadAPIRequest {
	return &AlibabaIhomeCtomContentImgUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIhomeCtomContentImgUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.ihome.ctom.content.img.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIhomeCtomContentImgUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MaterialDto Setter
// materialDTO
func (r *AlibabaIhomeCtomContentImgUploadAPIRequest) SetMaterialDto(_materialDto *UploadPicMaterialDto) error {
	r._materialDto = _materialDto
	r.Set("material_dto", _materialDto)
	return nil
}

// Get MaterialDto Getter
func (r AlibabaIhomeCtomContentImgUploadAPIRequest) GetMaterialDto() *UploadPicMaterialDto {
	return r._materialDto
}
