package paimai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaopaimainftorderinfoqueryAPIResponse 查询订单类型 API返回值
// taobao.paimai.nft.orderinfo.query
//
// 查询订单类型
type TaobaopaimainftorderinfoqueryAPIResponse struct {
	model.CommonResponse
	TaobaopaimainftorderinfoqueryAPIResponseModel
}

// TaobaopaimainftorderinfoqueryAPIResponseModel is 查询订单类型 成功返回结果
type TaobaopaimainftorderinfoqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"paimai_nft_orderinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单信息
	OrderInfoList []NftTradeOrderDto `json:"order_info_list,omitempty" xml:"order_info_list>nft_trade_order_dto,omitempty"`
	// 错误码
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 错误信息
	FailMessage string `json:"fail_message,omitempty" xml:"fail_message,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
