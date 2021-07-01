package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeGetAPIResponse
获取单笔交易的部分信息(性能高) API返回值
taobao.trade.get

获取单笔交易的部分信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong> */
type TaobaoTradeGetAPIResponse struct {
	model.CommonResponse
	TaobaoTradeGetAPIResponseModel
}

// TaobaoTradeGetAPIResponseModel is 获取单笔交易的部分信息(性能高) 成功返回结果
type TaobaoTradeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"trade_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息
	Trade *Trade `json:"trade,omitempty" xml:"trade,omitempty"`
}
