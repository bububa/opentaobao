package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
保存服务能力 API返回值 
alibaba.ssc.supplyplatform.serviceability.save

保存服务能力
*/
type AlibabaSscSupplyplatformServiceabilitySaveAPIResponse struct {
    model.CommonResponse
    AlibabaSscSupplyplatformServiceabilitySaveResponse
}

// 保存服务能力 成功返回结果
type AlibabaSscSupplyplatformServiceabilitySaveResponse struct {
    XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_serviceability_save_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaSscSupplyplatformServiceabilitySaveResult `json:"result,omitempty" xml:"result,omitempty"`
}
