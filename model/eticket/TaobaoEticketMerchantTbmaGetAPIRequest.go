package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoEticketMerchantTbmaGetAPIRequest
码商查询淘宝码接口 API请求
taobao.eticket.merchant.tbma.get

码商查询淘宝码接口 */
type TaobaoEticketMerchantTbmaGetAPIRequest struct {
	model.Params
	// 查询淘宝码请求
	_queryTbMaCallbackReq *QueryTbMaCallbackReq
}

// New
