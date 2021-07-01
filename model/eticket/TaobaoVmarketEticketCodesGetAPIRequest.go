package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketCodesGetAPIRequest
电子凭证码列表查询 API请求
taobao.vmarket.eticket.codes.get

查询某个订单的所有码的列表 */
type TaobaoVmarketEticketCodesGetAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
	// 码商ID
	_codemerchantId int64
}

// New
