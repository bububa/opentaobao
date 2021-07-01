package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelFutureSoftmodifyAPIRequest
未来酒店信息下发 API请求
taobao.xhotel.future.softmodify

未来酒店信息下发，包含PMS订单查询和自助入住 */
type TaobaoXhotelFutureSoftmodifyAPIRequest struct {
	model.Params
	// 超时时长，默认3s
	_expireTime int64
	// 淘宝订单号
	_tid int64
	// 外部订单号
	_outOrderId string
	// 酒店code
	_hotelCode string
	// 酒店Id
	_hid int64
	// 请求报文
	_context string
	// 操作类型
	_operateType string
	// 请求唯一标识值
	_requestId string
}

// New
