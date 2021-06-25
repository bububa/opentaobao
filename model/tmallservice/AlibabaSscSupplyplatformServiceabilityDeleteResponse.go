package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除服务能力 APIResponse
alibaba.ssc.supplyplatform.serviceability.delete

删除服务能力
*/
type AlibabaSscSupplyplatformServiceabilityDeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSscSupplyplatformServiceabilityDeleteResponse `json:"alibaba_ssc_supplyplatform_serviceability_delete_response,omitempty"`
}

type AlibabaSscSupplyplatformServiceabilityDeleteResponse struct {

    // 接口返回model
    Result   *AlibabaSscSupplyplatformServiceabilityDeleteResult `json:"result,omitempty"`

}