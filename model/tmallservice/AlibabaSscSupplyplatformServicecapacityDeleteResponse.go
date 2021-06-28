package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务容量删除 APIResponse
alibaba.ssc.supplyplatform.servicecapacity.delete

服务容量删除
*/
type AlibabaSscSupplyplatformServicecapacityDeleteAPIResponse struct {
    model.CommonResponse
    AlibabaSscSupplyplatformServicecapacityDeleteResponse
}

type AlibabaSscSupplyplatformServicecapacityDeleteResponse struct {
    XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_servicecapacity_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AlibabaSscSupplyplatformServicecapacityDeleteResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
