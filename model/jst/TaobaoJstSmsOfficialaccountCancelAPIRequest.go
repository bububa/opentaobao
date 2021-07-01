package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsOfficialaccountCancelAPIRequest
聚石塔取消公众号订购 API请求
taobao.jst.sms.officialaccount.cancel

聚石塔取消公众号订购 */
type TaobaoJstSmsOfficialaccountCancelAPIRequest struct {
	model.Params
	// 取消公众号订购请求
	_cancelOrderRequest *CancelOrderRequest
}

// New
