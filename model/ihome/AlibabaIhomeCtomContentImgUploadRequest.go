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
type AlibabaIhomeCtomContentImgUploadRequest struct {
    model.Params
    // materialDTO
    materialDto   *UploadPicMaterialDto
}

// 初始化AlibabaIhomeCtomContentImgUploadRequest对象
func NewAlibabaIhomeCtomContentImgUploadRequest() *AlibabaIhomeCtomContentImgUploadRequest{
    return &AlibabaIhomeCtomContentImgUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIhomeCtomContentImgUploadRequest) GetApiMethodName() string {
    return "alibaba.ihome.ctom.content.img.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIhomeCtomContentImgUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MaterialDto Setter
// materialDTO
func (r *AlibabaIhomeCtomContentImgUploadRequest) SetMaterialDto(materialDto *UploadPicMaterialDto) error {
    r.materialDto = materialDto
    r.Set("material_dto", materialDto)
    return nil
}

// MaterialDto Getter
func (r AlibabaIhomeCtomContentImgUploadRequest) GetMaterialDto() *UploadPicMaterialDto {
    return r.materialDto
}
