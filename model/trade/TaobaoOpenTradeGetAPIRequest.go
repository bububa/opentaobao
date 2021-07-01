package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenTradeGetAPIRequest
获取单笔交易的部分信息(商家应用使用) API请求
taobao.open.trade.get

获取单笔交易的部分信息</br>
1.入参fields中传入buyer_nick ，才能返回buyer_open_id */
type TaobaoOpenTradeGetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。
	_fields string
	// 交易编号
	_tid int64
}

// New
