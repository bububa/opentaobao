package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取合同类的服务工单信息 APIResponse
tmall.servicecenter.contracts.search

获取合同类的服务工单信息
*/
type TmallServicecenterContractsSearchAPIResponse struct {
    model.CommonResponse
    TmallServicecenterContractsSearchResponse
}

type TmallServicecenterContractsSearchResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_contracts_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 和同类服务包装类
    
    ServiceContractPacket   *ServiceContractPacket `json:"service_contract_packet,omitempty" xml:"service_contract_packet,omitempty"`

    
}
