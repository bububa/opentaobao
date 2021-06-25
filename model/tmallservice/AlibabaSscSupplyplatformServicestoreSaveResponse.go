package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
保存网点 APIResponse
alibaba.ssc.supplyplatform.servicestore.save

网点创建、修改
*/
type AlibabaSscSupplyplatformServicestoreSaveAPIResponse struct {
    model.CommonResponse
    Response *AlibabaSscSupplyplatformServicestoreSaveResponse `json:"alibaba_ssc_supplyplatform_servicestore_save_response,omitempty"`
}

type AlibabaSscSupplyplatformServicestoreSaveResponse struct {

    // 接口返回model
    Result   *AlibabaSscSupplyplatformServicestoreSaveResult `json:"result,omitempty"`

}
