package ihome

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实拍图投稿链路上传图片 APIRequest
alibaba.ihome.ctom.content.img.upload

实拍图投稿链路上传图片
*/
type AlibabaIhomeCtomContentImgUploadRequest struct {
    model.Params

    // materialDTO
    materialDto   *UploadPicMaterialDto 

}

func NewAlibabaIhomeCtomContentImgUploadRequest() *AlibabaIhomeCtomContentImgUploadRequest{
    return &AlibabaIhomeCtomContentImgUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIhomeCtomContentImgUploadRequest) GetApiMethodName() string {
    return "alibaba.ihome.ctom.content.img.upload"
}

func (r AlibabaIhomeCtomContentImgUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIhomeCtomContentImgUploadRequest) SetMaterialDto(materialDto *UploadPicMaterialDto) error {
    r.materialDto = materialDto
    r.Set("material_dto", materialDto)
    return nil
}

func (r AlibabaIhomeCtomContentImgUploadRequest) GetMaterialDto() *UploadPicMaterialDto {
    return r.materialDto
}

