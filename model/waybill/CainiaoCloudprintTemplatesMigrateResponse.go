package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云打印模板迁移接口 APIResponse
cainiao.cloudprint.templates.migrate

云打印模板迁移接口
*/
type CainiaoCloudprintTemplatesMigrateAPIResponse struct {
    model.CommonResponse
    CainiaoCloudprintTemplatesMigrateResponse
}

type CainiaoCloudprintTemplatesMigrateResponse struct {
    XMLName xml.Name `xml:"cainiao_cloudprint_templates_migrate_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
