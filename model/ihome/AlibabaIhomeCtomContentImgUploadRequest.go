package ihome

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实拍图投稿链路上传图片 API请求
alibaba.ihome.ctom.content.img.upload

实拍图投稿链路上传图片
*/
type AlibabaIhomeCtomContentImgUploadAPIRequest struct {
    model.Params
    // materialDTO
    _materialDto   *UploadPicMaterialDTO
}

// 初始化AlibabaIhomeCtomContentImgUploadAPIRequest对象
func NewAlibabaIhomeCtomContentImgUploadRequest() *AlibabaIhomeCtomContentImgUploadAPIRequest{
    return &AlibabaIhomeCtomContentImgUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIhomeCtomContentImgUploadAPIRequest) GetApiMethodName() string {
    return "alibaba.ihome.ctom.content.img.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIhomeCtomContentImgUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MaterialDto Setter
// materialDTO
func (r *AlibabaIhomeCtomContentImgUploadAPIRequest) SetMaterialDto(_materialDto *UploadPicMaterialDTO) error {
    r._materialDto = _materialDto
    r.Set("material_dto", _materialDto)
    return nil
}

// MaterialDto Getter
func (r AlibabaIhomeCtomContentImgUploadAPIRequest) GetMaterialDto() *UploadPicMaterialDTO {
    return r._materialDto
}
