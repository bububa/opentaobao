package ihome

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ihome图片上传 APIRequest
taobao.ihome.advancepic.upload

ihome 定制业务编辑器投稿素材上传
*/
type TaobaoIhomeAdvancepicUploadRequest struct {
    model.Params

    // 图片类
    materials   []AdvancePicMaterialDto 

}

func NewTaobaoIhomeAdvancepicUploadRequest() *TaobaoIhomeAdvancepicUploadRequest{
    return &TaobaoIhomeAdvancepicUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoIhomeAdvancepicUploadRequest) GetApiMethodName() string {
    return "taobao.ihome.advancepic.upload"
}

func (r TaobaoIhomeAdvancepicUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoIhomeAdvancepicUploadRequest) SetMaterials(materials []AdvancePicMaterialDto) error {
    r.materials = materials
    r.Set("materials", materials)
    return nil
}

func (r TaobaoIhomeAdvancepicUploadRequest) GetMaterials() []AdvancePicMaterialDto {
    return r.materials
}

