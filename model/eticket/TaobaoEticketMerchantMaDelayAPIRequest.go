package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoEticketMerchantMaDelayAPIRequest
凭证延期 API请求
taobao.eticket.merchant.ma.delay

订单延期 */
type TaobaoEticketMerchantMaDelayAPIRequest struct {
	model.Params
	// 业务类型
	_bizType int64
	// 延期时间
	_endDate string
	// 码
	_code string
	// 订单号
	_outerId string
	// 扩展
	_attributeMap string
	// 请求ID，调用方保证惟一
	_requestId string
}

// New
