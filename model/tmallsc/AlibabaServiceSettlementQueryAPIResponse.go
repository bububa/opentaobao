package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServiceSettlementQueryAPIResponse 服务平台结算单明细查询服务 API返回值
// alibaba.service.settlement.query
//
// 给服务商提供结算单明细查询功能
type AlibabaServiceSettlementQueryAPIResponse struct {
	model.CommonResponse
	AlibabaServiceSettlementQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaServiceSettlementQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaServiceSettlementQueryAPIResponseModel).Reset()
}

// AlibabaServiceSettlementQueryAPIResponseModel is 服务平台结算单明细查询服务 成功返回结果
type AlibabaServiceSettlementQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_service_settlement_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结算明细查询结果
	SettlementDetailQueryResult *FulfilplatformResult `json:"settlement_detail_query_result,omitempty" xml:"settlement_detail_query_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaServiceSettlementQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SettlementDetailQueryResult = nil
}

var poolAlibabaServiceSettlementQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaServiceSettlementQueryAPIResponse)
	},
}

// GetAlibabaServiceSettlementQueryAPIResponse 从 sync.Pool 获取 AlibabaServiceSettlementQueryAPIResponse
func GetAlibabaServiceSettlementQueryAPIResponse() *AlibabaServiceSettlementQueryAPIResponse {
	return poolAlibabaServiceSettlementQueryAPIResponse.Get().(*AlibabaServiceSettlementQueryAPIResponse)
}

// ReleaseAlibabaServiceSettlementQueryAPIResponse 将 AlibabaServiceSettlementQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaServiceSettlementQueryAPIResponse(v *AlibabaServiceSettlementQueryAPIResponse) {
	v.Reset()
	poolAlibabaServiceSettlementQueryAPIResponse.Put(v)
}
