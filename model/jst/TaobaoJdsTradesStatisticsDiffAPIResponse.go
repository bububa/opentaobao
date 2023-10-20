package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJdsTradesStatisticsDiffAPIResponse 订单全链路状态统计差异比较 API返回值
// taobao.jds.trades.statistics.diff
//
// 订单全链路状态统计差异比较
type TaobaoJdsTradesStatisticsDiffAPIResponse struct {
	model.CommonResponse
	TaobaoJdsTradesStatisticsDiffAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJdsTradesStatisticsDiffAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJdsTradesStatisticsDiffAPIResponseModel).Reset()
}

// TaobaoJdsTradesStatisticsDiffAPIResponseModel is 订单全链路状态统计差异比较 成功返回结果
type TaobaoJdsTradesStatisticsDiffAPIResponseModel struct {
	XMLName xml.Name `xml:"jds_trades_statistics_diff_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// pre_status比post_status多的tid列表
	Tids []int64 `json:"tids,omitempty" xml:"tids>int64,omitempty"`
	// 总记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJdsTradesStatisticsDiffAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Tids = m.Tids[:0]
	m.TotalResults = 0
}

var poolTaobaoJdsTradesStatisticsDiffAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJdsTradesStatisticsDiffAPIResponse)
	},
}

// GetTaobaoJdsTradesStatisticsDiffAPIResponse 从 sync.Pool 获取 TaobaoJdsTradesStatisticsDiffAPIResponse
func GetTaobaoJdsTradesStatisticsDiffAPIResponse() *TaobaoJdsTradesStatisticsDiffAPIResponse {
	return poolTaobaoJdsTradesStatisticsDiffAPIResponse.Get().(*TaobaoJdsTradesStatisticsDiffAPIResponse)
}

// ReleaseTaobaoJdsTradesStatisticsDiffAPIResponse 将 TaobaoJdsTradesStatisticsDiffAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJdsTradesStatisticsDiffAPIResponse(v *TaobaoJdsTradesStatisticsDiffAPIResponse) {
	v.Reset()
	poolTaobaoJdsTradesStatisticsDiffAPIResponse.Put(v)
}
