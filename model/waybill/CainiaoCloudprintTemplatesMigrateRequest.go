package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云打印模板迁移接口 APIRequest
cainiao.cloudprint.templates.migrate

云打印模板迁移接口
*/
type CainiaoCloudprintTemplatesMigrateRequest struct {
    model.Params

    // 标准电子面单模板的id
    tempalteId   int64 

    // 自定义区名称
    customAreaName   string 

    // 自定义区内容
    customAreaContent   string 

}

func NewCainiaoCloudprintTemplatesMigrateRequest() *CainiaoCloudprintTemplatesMigrateRequest{
    return &CainiaoCloudprintTemplatesMigrateRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoCloudprintTemplatesMigrateRequest) GetApiMethodName() string {
    return "cainiao.cloudprint.templates.migrate"
}

func (r CainiaoCloudprintTemplatesMigrateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoCloudprintTemplatesMigrateRequest) SetTempalteId(tempalteId int64) error {
    r.tempalteId = tempalteId
    r.Set("tempalte_id", tempalteId)
    return nil
}

func (r CainiaoCloudprintTemplatesMigrateRequest) GetTempalteId() int64 {
    return r.tempalteId
}

func (r *CainiaoCloudprintTemplatesMigrateRequest) SetCustomAreaName(customAreaName string) error {
    r.customAreaName = customAreaName
    r.Set("custom_area_name", customAreaName)
    return nil
}

func (r CainiaoCloudprintTemplatesMigrateRequest) GetCustomAreaName() string {
    return r.customAreaName
}

func (r *CainiaoCloudprintTemplatesMigrateRequest) SetCustomAreaContent(customAreaContent string) error {
    r.customAreaContent = customAreaContent
    r.Set("custom_area_content", customAreaContent)
    return nil
}

func (r CainiaoCloudprintTemplatesMigrateRequest) GetCustomAreaContent() string {
    return r.customAreaContent
}

