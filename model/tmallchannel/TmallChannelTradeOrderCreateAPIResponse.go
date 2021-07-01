package tmallchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallChannelTradeOrderCreateAPIResponse
创建渠道分销单 API返回值
tmall.channel.trade.order.create

创建渠道分销单 */
type TmallChannelTradeOrderCreateAPIResponse struct {
	model.CommonResponse
	TmallChannelTradeOrderCreateAPIResponseModel
}

// TmallChannelTradeOrderCreateAPIResponseModel is 创建渠道分销单 成功返回结果
type TmallChannelTradeOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_channel_trade_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 采购单号
	MainPurchaseOrderList []string `json:"main_purchase_order_list,omitempty" xml:"main_purchase_order_list>string,omitempty"`
}
