package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketManageNotifyAPIRequest
主动发起通知接口 API请求
taobao.vmarket.eticket.manage.notify

外部合作商家主动发起通知接口 */
type TaobaoVmarketEticketManageNotifyAPIRequest struct {
	model.Params
	// 订单编号
	_orderId int64
	// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
	_codemerchantId int64
	// 需要调用的通知方法，目前仅支持是send（发码）或resend（重新发码）
	_notifyMethod string
}

// New
