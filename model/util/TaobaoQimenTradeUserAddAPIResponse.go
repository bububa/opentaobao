package util

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTradeUserAddAPIResponse 添加奇门订单链路用户 API返回值
// taobao.qimen.trade.user.add
//
// 添加奇门订单链路用户
type TaobaoQimenTradeUserAddAPIResponse struct {
	model.CommonResponse
	TaobaoQimenTradeUserAddAPIResponseModel
}

// TaobaoQimenTradeUserAddAPIResponseModel is 添加奇门订单链路用户 成功返回结果
type TaobaoQimenTradeUserAddAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_trade_user_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 卖家备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
}
