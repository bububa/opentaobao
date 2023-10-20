package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJdsTradesStatisticsGetAPIResponse 获取订单数量统计结果 API返回值
// taobao.jds.trades.statistics.get
//
// 获取订单数量统计结果
type TaobaoJdsTradesStatisticsGetAPIResponse struct {
	model.CommonResponse
	TaobaoJdsTradesStatisticsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJdsTradesStatisticsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJdsTradesStatisticsGetAPIResponseModel).Reset()
}

// TaobaoJdsTradesStatisticsGetAPIResponseModel is 获取订单数量统计结果 成功返回结果
type TaobaoJdsTradesStatisticsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"jds_trades_statistics_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单状态计数值
	StatusInfos []TradeStat `json:"status_infos,omitempty" xml:"status_infos>trade_stat,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJdsTradesStatisticsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StatusInfos = m.StatusInfos[:0]
}

var poolTaobaoJdsTradesStatisticsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJdsTradesStatisticsGetAPIResponse)
	},
}

// GetTaobaoJdsTradesStatisticsGetAPIResponse 从 sync.Pool 获取 TaobaoJdsTradesStatisticsGetAPIResponse
func GetTaobaoJdsTradesStatisticsGetAPIResponse() *TaobaoJdsTradesStatisticsGetAPIResponse {
	return poolTaobaoJdsTradesStatisticsGetAPIResponse.Get().(*TaobaoJdsTradesStatisticsGetAPIResponse)
}

// ReleaseTaobaoJdsTradesStatisticsGetAPIResponse 将 TaobaoJdsTradesStatisticsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJdsTradesStatisticsGetAPIResponse(v *TaobaoJdsTradesStatisticsGetAPIResponse) {
	v.Reset()
	poolTaobaoJdsTradesStatisticsGetAPIResponse.Put(v)
}
