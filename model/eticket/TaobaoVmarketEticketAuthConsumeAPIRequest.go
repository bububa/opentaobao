package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketAuthConsumeAPIRequest
核销放行的核销接口 API请求
taobao.vmarket.eticket.auth.consume

针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销接口 */
type TaobaoVmarketEticketAuthConsumeAPIRequest struct {
	model.Params
	// 核销的码，只支持单个码，多个码核销需要多次调用
	_verifyCode string
	// 核销份数
	_consumeNum int64
	// 核销方的ID，如果是普通码商必须传入机具ID,如果是私有码商家（即原有的信任商家）可默认传入私有码商ID
	_operatorid string
	// 自定义核销流水号，需要小于等于100个字符(a-zA-Z0-9_)
	_serialNum string
	// 网点ID,网点授权核销时，必须传入；其他核销方式可不传
	_storeid string
}

// New
