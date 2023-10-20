package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaojdstradesstatisticsdiffAPIResponse 订单全链路状态统计差异比较 API返回值
// taobao.jds.trades.statistics.diff
//
// 订单全链路状态统计差异比较
type TaobaojdstradesstatisticsdiffAPIResponse struct {
	model.CommonResponse
	TaobaojdstradesstatisticsdiffAPIResponseModel
}

// TaobaojdstradesstatisticsdiffAPIResponseModel is 订单全链路状态统计差异比较 成功返回结果
type TaobaojdstradesstatisticsdiffAPIResponseModel struct {
	XMLName xml.Name `xml:"jds_trades_statistics_diff_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// pre_status比post_status多的tid列表
	Tids []int64 `json:"tids,omitempty" xml:"tids>int64,omitempty"`
	// 总记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
