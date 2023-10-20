package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServiceSettlementBillinfoQueryAPIResponse 查询工单结算信息 API返回值
// tmall.service.settlement.billinfo.query
//
// 提供给服务商查询工单结算信息，包含结算的分成金额以及结算的收款明细，平台抽佣比例。用于服务商进行对账
type TmallServiceSettlementBillinfoQueryAPIResponse struct {
	model.CommonResponse
	TmallServiceSettlementBillinfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServiceSettlementBillinfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServiceSettlementBillinfoQueryAPIResponseModel).Reset()
}

// TmallServiceSettlementBillinfoQueryAPIResponseModel is 查询工单结算信息 成功返回结果
type TmallServiceSettlementBillinfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_service_settlement_billinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口响应
	Result *TmallServiceSettlementBillinfoQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServiceSettlementBillinfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServiceSettlementBillinfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServiceSettlementBillinfoQueryAPIResponse)
	},
}

// GetTmallServiceSettlementBillinfoQueryAPIResponse 从 sync.Pool 获取 TmallServiceSettlementBillinfoQueryAPIResponse
func GetTmallServiceSettlementBillinfoQueryAPIResponse() *TmallServiceSettlementBillinfoQueryAPIResponse {
	return poolTmallServiceSettlementBillinfoQueryAPIResponse.Get().(*TmallServiceSettlementBillinfoQueryAPIResponse)
}

// ReleaseTmallServiceSettlementBillinfoQueryAPIResponse 将 TmallServiceSettlementBillinfoQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallServiceSettlementBillinfoQueryAPIResponse(v *TmallServiceSettlementBillinfoQueryAPIResponse) {
	v.Reset()
	poolTmallServiceSettlementBillinfoQueryAPIResponse.Put(v)
}
