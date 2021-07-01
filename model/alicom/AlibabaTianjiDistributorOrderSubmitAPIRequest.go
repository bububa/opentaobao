package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTianjiDistributorOrderSubmitAPIRequest
分销商提交受理订单 API请求
alibaba.tianji.distributor.order.submit

分销商提交受理订单，如合约订购、充值受理等 */
type AlibabaTianjiDistributorOrderSubmitAPIRequest struct {
	model.Params
	// 商品编码，如手机串号
	_itemSerialNo string
	// 淘宝交易订单号
	_orderNo string
	// 供应商产品编码，如SIM卡号
	_productSerialNo string
}

// New
