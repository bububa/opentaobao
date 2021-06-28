package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务容量删除 APIResponse
alibaba.ssc.supplyplatform.servicecapacity.delete

服务容量删除
*/
type AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSscSupplyplatformServicecapacityDeleteResponse `json:"alibaba_ssc_supplyplatform_servicecapacity_delete_response,omitempty"` 
    AlibabaSscSupplyplatformServicecapacityDeleteResponse
}

/* model for simplify = false
type AlibabaSscSupplyplatformServicecapacityDeleteResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaSscSupplyplatformServicecapacityDeleteResult  *AlibabaSscSupplyplatformServicecapacityDeleteResult `json:"alibaba_ssc_supplyplatform_servicecapacity_delete_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaSscSupplyplatformServicecapacityDeleteResponse struct {

    // 接口返回model
    Result   *AlibabaSscSupplyplatformServicecapacityDeleteResult `json:"result,omitempty"`

}
