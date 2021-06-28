package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取商家使用的标准模板 APIResponse
cainiao.cloudprint.isvtemplates.get

获取商家使用的标准模板
*/
type CainiaoCloudprintIsvtemplatesGetAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoCloudprintIsvtemplatesGetResponse `json:"cainiao_cloudprint_isvtemplates_get_response,omitempty"` 
    CainiaoCloudprintIsvtemplatesGetResponse
}

/* model for simplify = false
type CainiaoCloudprintIsvtemplatesGetResponse struct {

    // result
    
    Result  *struct {
        CloudPrintBaseResult  *CloudPrintBaseResult `json:"cloud_print_base_result,omitempty"`
    } `json:"result,omitempty"`
    

    // result
    
    Result  *struct {
        CloudPrintBaseResult  *CloudPrintBaseResult `json:"cloud_print_base_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type CainiaoCloudprintIsvtemplatesGetResponse struct {

    // result
    Result   *CloudPrintBaseResult `json:"result,omitempty"`

    // result
    Result   *CloudPrintBaseResult `json:"result,omitempty"`

}
