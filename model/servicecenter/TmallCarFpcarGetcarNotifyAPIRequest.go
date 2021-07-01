package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarFpcarGetcarNotifyAPIRequest
门店通知用户提车 API请求
tmall.car.fpcar.getcar.notify

提供给外部(大搜或其它合作方)的接口-门店通知用户提车 */
type TmallCarFpcarGetcarNotifyAPIRequest struct {
	model.Params
	// 商品宝贝id
	_itemId int64
	// 订单id
	_orderId int64
	// 卖家id
	_sellerId int64
}

// New
