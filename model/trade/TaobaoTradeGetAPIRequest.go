package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeGetAPIRequest
获取单笔交易的部分信息(性能高) API请求
taobao.trade.get

获取单笔交易的部分信息
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=tradeapi" target="_blank">点击查看更多交易API说明</a></strong> */
type TaobaoTradeGetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
	_fields string
	// 交易编号
	_tid int64
}

// New
