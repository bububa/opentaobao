package rail

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripRailTradeCloseticketAPIResponse 出票失败关单接口 API返回值
// alitrip.rail.trade.closeticket
//
// 出票成功回调接口
type AlitripRailTradeCloseticketAPIResponse struct {
	model.CommonResponse
	AlitripRailTradeCloseticketAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripRailTradeCloseticketAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripRailTradeCloseticketAPIResponseModel).Reset()
}

// AlitripRailTradeCloseticketAPIResponseModel is 出票失败关单接口 成功返回结果
type AlitripRailTradeCloseticketAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_rail_trade_closeticket_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *AlitripRailTradeCloseticketResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripRailTradeCloseticketAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripRailTradeCloseticketAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripRailTradeCloseticketAPIResponse)
	},
}

// GetAlitripRailTradeCloseticketAPIResponse 从 sync.Pool 获取 AlitripRailTradeCloseticketAPIResponse
func GetAlitripRailTradeCloseticketAPIResponse() *AlitripRailTradeCloseticketAPIResponse {
	return poolAlitripRailTradeCloseticketAPIResponse.Get().(*AlitripRailTradeCloseticketAPIResponse)
}

// ReleaseAlitripRailTradeCloseticketAPIResponse 将 AlitripRailTradeCloseticketAPIResponse 保存到 sync.Pool
func ReleaseAlitripRailTradeCloseticketAPIResponse(v *AlitripRailTradeCloseticketAPIResponse) {
	v.Reset()
	poolAlitripRailTradeCloseticketAPIResponse.Put(v)
}
