package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取合同类的服务工单信息 APIResponse
tmall.servicecenter.contracts.search

获取合同类的服务工单信息
*/
type TmallServicecenterContractsSearchAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterContractsSearchResponse `json:"tmall_servicecenter_contracts_search_response,omitempty"`
}

type TmallServicecenterContractsSearchResponse struct {

    // 和同类服务包装类
    ServiceContractPacket   *ServiceContractPacket `json:"service_contract_packet,omitempty"`

}
