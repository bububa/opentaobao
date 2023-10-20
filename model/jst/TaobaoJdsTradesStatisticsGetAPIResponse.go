package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaojdstradesstatisticsgetAPIResponse 获取订单数量统计结果 API返回值
// taobao.jds.trades.statistics.get
//
// 获取订单数量统计结果
type TaobaojdstradesstatisticsgetAPIResponse struct {
	model.CommonResponse
	TaobaojdstradesstatisticsgetAPIResponseModel
}

// TaobaojdstradesstatisticsgetAPIResponseModel is 获取订单数量统计结果 成功返回结果
type TaobaojdstradesstatisticsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"jds_trades_statistics_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单状态计数值
	StatusInfos []TradeStat `json:"status_infos,omitempty" xml:"status_infos>trade_stat,omitempty"`
}
