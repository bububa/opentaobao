package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商绑定工人 APIResponse
alibaba.ssc.supplyplatform.serviceworker.save

服务商将上传工人与服务商自己建立关系，需要将工人的服务区域和住址回传
*/
type AlibabaSscSupplyplatformServiceworkerSaveAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSscSupplyplatformServiceworkerSaveResponse `json:"alibaba_ssc_supplyplatform_serviceworker_save_response,omitempty"` 
    AlibabaSscSupplyplatformServiceworkerSaveResponse
}

/* model for simplify = false
type AlibabaSscSupplyplatformServiceworkerSaveResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaSscSupplyplatformServiceworkerSaveResult  *AlibabaSscSupplyplatformServiceworkerSaveResult `json:"alibaba_ssc_supplyplatform_serviceworker_save_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaSscSupplyplatformServiceworkerSaveResponse struct {

    // 接口返回model
    Result   *AlibabaSscSupplyplatformServiceworkerSaveResult `json:"result,omitempty"`

}
