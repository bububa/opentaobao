package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
保存服务容量 APIResponse
alibaba.ssc.supplyplatform.servicecapacity.save

保存服务容量
*/
type AlibabaSscSupplyplatformServicecapacitySaveAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSscSupplyplatformServicecapacitySaveResponse `json:"alibaba_ssc_supplyplatform_servicecapacity_save_response,omitempty"`
}

type AlibabaSscSupplyplatformServicecapacitySaveResponse struct {

    // 接口返回model
    Result   *AlibabaSscSupplyplatformServicecapacitySaveResult `json:"result,omitempty"`

}
