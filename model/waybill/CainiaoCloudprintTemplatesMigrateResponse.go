package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
云打印模板迁移接口 APIResponse
cainiao.cloudprint.templates.migrate

云打印模板迁移接口
*/
type CainiaoCloudprintTemplatesMigrateAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoCloudprintTemplatesMigrateResponse `json:"cainiao_cloudprint_templates_migrate_response,omitempty"` 
    CainiaoCloudprintTemplatesMigrateResponse
}

/* model for simplify = false
type CainiaoCloudprintTemplatesMigrateResponse struct {

    // result
    
    Result  *struct {
        CloudPrintBaseResult  *CloudPrintBaseResult `json:"cloud_print_base_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type CainiaoCloudprintTemplatesMigrateResponse struct {

    // result
    Result   *CloudPrintBaseResult `json:"result,omitempty"`

}
