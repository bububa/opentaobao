package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
保存服务能力 APIResponse
alibaba.ssc.supplyplatform.serviceability.save

保存服务能力
*/
type AlibabaSscSupplyplatformServiceabilitySaveAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSscSupplyplatformServiceabilitySaveResponse `json:"alibaba_ssc_supplyplatform_serviceability_save_response,omitempty"` 
    AlibabaSscSupplyplatformServiceabilitySaveResponse
}

/* model for simplify = false
type AlibabaSscSupplyplatformServiceabilitySaveResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaSscSupplyplatformServiceabilitySaveResult  *AlibabaSscSupplyplatformServiceabilitySaveResult `json:"alibaba_ssc_supplyplatform_serviceability_save_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaSscSupplyplatformServiceabilitySaveResponse struct {

    // 接口返回model
    Result   *AlibabaSscSupplyplatformServiceabilitySaveResult `json:"result,omitempty"`

}
