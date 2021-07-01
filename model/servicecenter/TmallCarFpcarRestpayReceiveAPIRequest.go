package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarFpcarRestpayReceiveAPIRequest
门店线下已收尾款 API请求
tmall.car.fpcar.restpay.receive

提供给外部(大搜或其它合作方)的接口-门店线下已收尾款(不执行分佣) */
type TmallCarFpcarRestpayReceiveAPIRequest struct {
	model.Params
	// 卖家id
	_sellerId int64
	// 订单id
	_orderId int64
	// 商品宝贝id
	_itemId int64
}

// New
