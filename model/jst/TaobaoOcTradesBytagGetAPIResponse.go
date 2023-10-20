package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaooctradesbytaggetAPIResponse 标签查询订单 API返回值
// taobao.oc.trades.bytag.get
//
// 根据标签查询订单编号
type TaobaooctradesbytaggetAPIResponse struct {
	model.CommonResponse
	TaobaooctradesbytaggetAPIResponseModel
}

// TaobaooctradesbytaggetAPIResponseModel is 标签查询订单 成功返回结果
type TaobaooctradesbytaggetAPIResponseModel struct {
	XMLName xml.Name `xml:"oc_trades_bytag_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 打了该标签的订单编号列表
	Tids []int64 `json:"tids,omitempty" xml:"tids>int64,omitempty"`
	// 总数
	Totals int64 `json:"totals,omitempty" xml:"totals,omitempty"`
}
