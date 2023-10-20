package rail

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripRailTradeIssueticketAPIResponse 德铁出票成功接口 API返回值
// alitrip.rail.trade.issueticket
//
// 出票成功回调接口
type AlitripRailTradeIssueticketAPIResponse struct {
	model.CommonResponse
	AlitripRailTradeIssueticketAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripRailTradeIssueticketAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripRailTradeIssueticketAPIResponseModel).Reset()
}

// AlitripRailTradeIssueticketAPIResponseModel is 德铁出票成功接口 成功返回结果
type AlitripRailTradeIssueticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_rail_trade_issueticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *AlitripRailTradeIssueticketResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripRailTradeIssueticketAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripRailTradeIssueticketAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripRailTradeIssueticketAPIResponse)
	},
}

// GetAlitripRailTradeIssueticketAPIResponse 从 sync.Pool 获取 AlitripRailTradeIssueticketAPIResponse
func GetAlitripRailTradeIssueticketAPIResponse() *AlitripRailTradeIssueticketAPIResponse {
	return poolAlitripRailTradeIssueticketAPIResponse.Get().(*AlitripRailTradeIssueticketAPIResponse)
}

// ReleaseAlitripRailTradeIssueticketAPIResponse 将 AlitripRailTradeIssueticketAPIResponse 保存到 sync.Pool
func ReleaseAlitripRailTradeIssueticketAPIResponse(v *AlitripRailTradeIssueticketAPIResponse) {
	v.Reset()
	poolAlitripRailTradeIssueticketAPIResponse.Put(v)
}
