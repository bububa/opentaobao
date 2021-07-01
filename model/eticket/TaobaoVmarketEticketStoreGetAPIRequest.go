package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketStoreGetAPIRequest
获取电子凭证预约门店信息 API请求
taobao.vmarket.eticket.store.get

用于给外部商家查询电子凭证预约门店信息 */
type TaobaoVmarketEticketStoreGetAPIRequest struct {
	model.Params
	// 订单ID
	_orderId int64
}

// New
