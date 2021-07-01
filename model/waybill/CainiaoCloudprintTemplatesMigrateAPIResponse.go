package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云打印模板迁移接口 API返回值 
cainiao.cloudprint.templates.migrate

云打印模板迁移接口
*/
type CainiaoCloudprintTemplatesMigrateAPIResponse struct {
    model.CommonResponse
    CainiaoCloudprintTemplatesMigrateAPIResponseModel
}

// 云打印模板迁移接口 成功返回结果
type CainiaoCloudprintTemplatesMigrateAPIResponseModel struct {
    XMLName xml.Name `xml:"cainiao_cloudprint_templates_migrate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
