package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云打印模板迁移接口 API请求
cainiao.cloudprint.templates.migrate

云打印模板迁移接口
*/
type CainiaoCloudprintTemplatesMigrateAPIRequest struct {
    model.Params
    // 标准电子面单模板的id
    _tempalteId   int64
    // 自定义区名称
    _customAreaName   string
    // 自定义区内容
    _customAreaContent   string
}

// 初始化CainiaoCloudprintTemplatesMigrateAPIRequest对象
func NewCainiaoCloudprintTemplatesMigrateRequest() *CainiaoCloudprintTemplatesMigrateAPIRequest{
    return &CainiaoCloudprintTemplatesMigrateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoCloudprintTemplatesMigrateAPIRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.templates.migrate"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoCloudprintTemplatesMigrateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TempalteId Setter
// 标准电子面单模板的id
func (r *CainiaoCloudprintTemplatesMigrateAPIRequest) SetTempalteId(_tempalteId int64) error {
    r._tempalteId = _tempalteId
    r.Set("tempalte_id", _tempalteId)
    return nil
}

// TempalteId Getter
func (r CainiaoCloudprintTemplatesMigrateAPIRequest) GetTempalteId() int64 {
    return r._tempalteId
}
// CustomAreaName Setter
// 自定义区名称
func (r *CainiaoCloudprintTemplatesMigrateAPIRequest) SetCustomAreaName(_customAreaName string) error {
    r._customAreaName = _customAreaName
    r.Set("custom_area_name", _customAreaName)
    return nil
}

// CustomAreaName Getter
func (r CainiaoCloudprintTemplatesMigrateAPIRequest) GetCustomAreaName() string {
    return r._customAreaName
}
// CustomAreaContent Setter
// 自定义区内容
func (r *CainiaoCloudprintTemplatesMigrateAPIRequest) SetCustomAreaContent(_customAreaContent string) error {
    r._customAreaContent = _customAreaContent
    r.Set("custom_area_content", _customAreaContent)
    return nil
}

// CustomAreaContent Getter
func (r CainiaoCloudprintTemplatesMigrateAPIRequest) GetCustomAreaContent() string {
    return r._customAreaContent
}
