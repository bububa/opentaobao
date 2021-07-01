package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自定义区内容更新 API请求
cainiao.cloudprint.customarea.update

自定义区内容更新
*/
type CainiaoCloudprintCustomareaUpdateAPIRequest struct {
    model.Params
    // 自定义区id（不可修改）
    _customAreaId   int64
    // 自定义区名称（可修改）
    _customAreaName   string
    // 自定义区内容（可修改）
    _customAreaContent   string
}

// 初始化CainiaoCloudprintCustomareaUpdateAPIRequest对象
func NewCainiaoCloudprintCustomareaUpdateRequest() *CainiaoCloudprintCustomareaUpdateAPIRequest{
    return &CainiaoCloudprintCustomareaUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintCustomareaUpdateAPIRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.customarea.update"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintCustomareaUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CustomAreaId Setter
// 自定义区id（不可修改）
func (r *CainiaoCloudprintCustomareaUpdateAPIRequest) SetCustomAreaId(_customAreaId int64) error {
    r._customAreaId = _customAreaId
    r.Set("custom_area_id", _customAreaId)
    return nil
}

// CustomAreaId Getter
func (r CainiaoCloudprintCustomareaUpdateAPIRequest) GetCustomAreaId() int64 {
    return r._customAreaId
}
// CustomAreaName Setter
// 自定义区名称（可修改）
func (r *CainiaoCloudprintCustomareaUpdateAPIRequest) SetCustomAreaName(_customAreaName string) error {
    r._customAreaName = _customAreaName
    r.Set("custom_area_name", _customAreaName)
    return nil
}

// CustomAreaName Getter
func (r CainiaoCloudprintCustomareaUpdateAPIRequest) GetCustomAreaName() string {
    return r._customAreaName
}
// CustomAreaContent Setter
// 自定义区内容（可修改）
func (r *CainiaoCloudprintCustomareaUpdateAPIRequest) SetCustomAreaContent(_customAreaContent string) error {
    r._customAreaContent = _customAreaContent
    r.Set("custom_area_content", _customAreaContent)
    return nil
}

// CustomAreaContent Getter
func (r CainiaoCloudprintCustomareaUpdateAPIRequest) GetCustomAreaContent() string {
    return r._customAreaContent
}
