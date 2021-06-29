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
type CainiaoCloudprintCustomareaUpdateRequest struct {
    model.Params
    // 自定义区id（不可修改）
    customAreaId   int64
    // 自定义区名称（可修改）
    customAreaName   string
    // 自定义区内容（可修改）
    customAreaContent   string
}

// 初始化CainiaoCloudprintCustomareaUpdateRequest对象
func NewCainiaoCloudprintCustomareaUpdateRequest() *CainiaoCloudprintCustomareaUpdateRequest{
    return &CainiaoCloudprintCustomareaUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintCustomareaUpdateRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.customarea.update"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintCustomareaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CustomAreaId Setter
// 自定义区id（不可修改）
func (r *CainiaoCloudprintCustomareaUpdateRequest) SetCustomAreaId(customAreaId int64) error {
    r.customAreaId = customAreaId
    r.Set("custom_area_id", customAreaId)
    return nil
}

// CustomAreaId Getter
func (r CainiaoCloudprintCustomareaUpdateRequest) GetCustomAreaId() int64 {
    return r.customAreaId
}
// CustomAreaName Setter
// 自定义区名称（可修改）
func (r *CainiaoCloudprintCustomareaUpdateRequest) SetCustomAreaName(customAreaName string) error {
    r.customAreaName = customAreaName
    r.Set("custom_area_name", customAreaName)
    return nil
}

// CustomAreaName Getter
func (r CainiaoCloudprintCustomareaUpdateRequest) GetCustomAreaName() string {
    return r.customAreaName
}
// CustomAreaContent Setter
// 自定义区内容（可修改）
func (r *CainiaoCloudprintCustomareaUpdateRequest) SetCustomAreaContent(customAreaContent string) error {
    r.customAreaContent = customAreaContent
    r.Set("custom_area_content", customAreaContent)
    return nil
}

// CustomAreaContent Getter
func (r CainiaoCloudprintCustomareaUpdateRequest) GetCustomAreaContent() string {
    return r.customAreaContent
}
