package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterContractsSearchAPIResponse
获取合同类的服务工单信息 API返回值
tmall.servicecenter.contracts.search

获取合同类的服务工单信息 */
type TmallServicecenterContractsSearchAPIResponse struct {
	model.CommonResponse
	TmallServicecenterContractsSearchAPIResponseModel
}

// TmallServicecenterContractsSearchAPIResponseModel is 获取合同类的服务工单信息 成功返回结果
type TmallServicecenterContractsSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_contracts_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和同类服务包装类
	ServiceContractPacket *ServiceContractPacket `json:"service_contract_packet,omitempty" xml:"service_contract_packet,omitempty"`
}
