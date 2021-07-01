package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinTccTradeIdentityGetAPIRequest
运营商获得用户身份信息 API请求
alibaba.aliqin.tcc.trade.identity.get

天猫网厅运营商官方旗舰店获取用户身份信息 */
type AlibabaAliqinTccTradeIdentityGetAPIRequest struct {
	model.Params
	// 订单编号
	_bizOrderId int64
	// 店铺名称
	_sellerNick string
}

// New
