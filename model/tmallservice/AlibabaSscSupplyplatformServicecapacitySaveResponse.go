package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
保存服务容量 APIResponse
alibaba.ssc.supplyplatform.servicecapacity.save

保存服务容量
*/
type AlibabaSscSupplyplatformServicecapacitySaveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ssc_supplyplatform_servicecapacity_save_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AlibabaSscSupplyplatformServicecapacitySaveResult `json:"result,omitempty" xml:"