package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
isv资源查询 APIResponse
cainiao.cloudprint.isv.resources.get

isv资源查询，包括isv模板、打印项、预设的自定义区等
*/
type CainiaoCloudprintIsvResourcesGetAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoCloudprintIsvResourcesGetResponse `json:"cainiao_cloudprint_isv_resources_get_response,omitempty"` 
    CainiaoCloudprintIsvResourcesGetResponse
}

/* model for simplify = false
type CainiaoCloudprintIsvResourcesGetResponse struct {

    // result
    
    Result  *struct {
        CloudPrintBaseResult  *CloudPrintBaseResult `json:"cloud_print_base_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type CainiaoCloudprintIsvResourcesGetResponse struct {

    // result
    Result   *CloudPrintBaseResult `json:"result,omitempty"`

}
