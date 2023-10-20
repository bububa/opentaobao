package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOnsiteTradeQueryrefundAPIResponse 退款查询 API返回值
// alibaba.mos.onsite.trade.queryrefund
//
// 商户可使用该接口查询退款请求是否执行成功。
type AlibabaMosOnsiteTradeQueryrefundAPIResponse struct {
	model.CommonResponse
	AlibabaMosOnsiteTradeQueryrefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosOnsiteTradeQueryrefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosOnsiteTradeQueryrefundAPIResponseModel).Reset()
}

// AlibabaMosOnsiteTradeQueryrefundAPIResponseModel is 退款查询 成功返回结果
type AlibabaMosOnsiteTradeQueryrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_onsite_trade_queryrefund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaMosOnsiteTradeQueryrefundResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosOnsiteTradeQueryrefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosOnsiteTradeQueryrefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosOnsiteTradeQueryrefundAPIResponse)
	},
}

// GetAlibabaMosOnsiteTradeQueryrefundAPIResponse 从 sync.Pool 获取 AlibabaMosOnsiteTradeQueryrefundAPIResponse
func GetAlibabaMosOnsiteTradeQueryrefundAPIResponse() *AlibabaMosOnsiteTradeQueryrefundAPIResponse {
	return poolAlibabaMosOnsiteTradeQueryrefundAPIResponse.Get().(*AlibabaMosOnsiteTradeQueryrefundAPIResponse)
}

// ReleaseAlibabaMosOnsiteTradeQueryrefundAPIResponse 将 AlibabaMosOnsiteTradeQueryrefundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosOnsiteTradeQueryrefundAPIResponse(v *AlibabaMosOnsiteTradeQueryrefundAPIResponse) {
	v.Reset()
	poolAlibabaMosOnsiteTradeQueryrefundAPIResponse.Put(v)
}
