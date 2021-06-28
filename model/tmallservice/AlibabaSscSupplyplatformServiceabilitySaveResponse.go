package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
保存服务能力 APIResponse
alibaba.ssc.supplyplatform.serviceability.save

保存服务能力
*/
type AlibabaSscSupplyplatformServiceabilitySaveAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ssc_supplyplatform_serviceability_save_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AlibabaSscSupplyplatformServiceabilitySaveResult `json:"result,omitempty" xml:"