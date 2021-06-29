package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据天猫id查询门店信息 APIResponse
alibaba.ssc.servicecenter.servicestore.query

根据天猫id查询门店信息
*/
type AlibabaSscServicecenterServicestoreQueryAPIResponse struct {
    model.CommonResponse
    AlibabaSscServicecenterServicestoreQueryResponse
}

type AlibabaSscServicecenterServicestoreQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_ssc_servicecenter_servicestore_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaSscServicecenterServicestoreQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
