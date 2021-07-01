package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsOfficialaccountOrderAPIRequest
聚石塔订购公众号 API请求
taobao.jst.sms.officialaccount.order

聚石塔订购公众号接口 */
type TaobaoJstSmsOfficialaccountOrderAPIRequest struct {
	model.Params
	// 聚石塔公众号订购
	_orderOfficialAccountRequest *OrderOfficialAccountRequest
}

// New
