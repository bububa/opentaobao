package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
网点下线 APIResponse
alibaba.ssc.supplyplatform.servicestore.offline

网点下线功能
*/
type AlibabaSscSupplyplatformServicestoreOfflineAPIResponse struct {
    model.CommonResponse
    AlibabaSscSupplyplatformServicestoreOfflineResponse
}

type AlibabaSscSupplyplatformServicestoreOfflineResponse struct {
    XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_servicestore_offline_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaSscSupplyplatformServicestoreOfflineResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
