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
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_cloudprint_templates_migrate_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *CloudPrintBaseResult `json:"result,omitempty" xml:"