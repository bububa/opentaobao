package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取海王用户权限包 APIResponse
alibaba.seaking.servicepack

获取海王用户权限包
*/
type AlibabaSeakingServicepackAPIResponse struct {
    model.CommonResponse
    AlibabaSeakingServicepackResponse
}

type AlibabaSeakingServicepackResponse struct {
    XMLName xml.Name `xml:"alibaba_seaking_servicepack_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 权限包列表
    
    ServicePackList   []AlibabaSeakingServicepackResult `json:"service_pack_list,omitempty" xml:"service_pack_list>alibaba_seaking_servicepack_result,omitempty"`
    
    
}
