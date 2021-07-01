package ihome

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ihome图片上传 API请求
taobao.ihome.advancepic.upload

ihome 定制业务编辑器投稿素材上传
*/
type TaobaoIhomeAdvancepicUploadAPIRequest struct {
    model.Params
    // 图片类
    _materials   []AdvancePicMaterialDto
}

// 初始化TaobaoIhomeAdvancepicUploadAPIRequest对象
func NewTaobaoIhomeAdvancepicUploadRequest() *TaobaoIhomeAdvancepicUploadAPIRequest{
    return &TaobaoIhomeAdvancepicUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoIhomeAdvancepicUploadAPIRequest) GetApiMethodName() string {
    return "taobao.ihome.advancepic.upload"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoIhomeAdvancepicUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Materials Setter
// 图片类
func (r *TaobaoIhomeAdvancepicUploadAPIRequest) SetMaterials(_materials []AdvancePicMaterialDto) error {
    r._materials = _materials
    r.Set("materials", _materials)
    return nil
}

// Materials Getter
func (r TaobaoIhomeAdvancepicUploadAPIRequest) GetMaterials() []AdvancePicMaterialDto {
    return r._materials
}
