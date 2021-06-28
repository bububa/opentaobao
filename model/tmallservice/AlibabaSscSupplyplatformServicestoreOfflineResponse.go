package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
网点下线 APIResponse
alibaba.ssc.supplyplatform.servicestore.offline

网点下线功能
*/
type AlibabaSscSupplyplatformServicestoreOfflineAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSscSupplyplatformServicestoreOfflineResponse `json:"alibaba_ssc_supplyplatform_servicestore_offline_response,omitempty"` 
    AlibabaSscSupplyplatformServicestoreOfflineResponse
}

/* model for simplify = false
type AlibabaSscSupplyplatformServicestoreOfflineResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaSscSupplyplatformServicestoreOfflineResult  *AlibabaSscSupplyplatformServicestoreOfflineResult `json:"alibaba_ssc_supplyplatform_servicestore_offline_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaSscSupplyplatformServicestoreOfflineResponse struct {

    // 接口返回model
    Result   *AlibabaSscSupplyplatformServicestoreOfflineResult `json:"result,omitempty"`

}
