package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取商家的自定义区模板信息 APIResponse
cainiao.cloudprint.customares.get

供isv使用，获取商家的自定义区的模板信息
*/
type CainiaoCloudprintCustomaresGetAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoCloudprintCustomaresGetResponse `json:"cainiao_cloudprint_customares_get_response,omitempty"` 
    CainiaoCloudprintCustomaresGetResponse
}

/* model for simplify = false
type CainiaoCloudprintCustomaresGetResponse struct {

    // 结果
    
    Result  *struct {
        CloudPrintBaseResult  *CloudPrintBaseResult `json:"cloud_print_base_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type CainiaoCloudprintCustomaresGetResponse struct {

    // 结果
    Result   *CloudPrintBaseResult `json:"result,omitempty"`

}
