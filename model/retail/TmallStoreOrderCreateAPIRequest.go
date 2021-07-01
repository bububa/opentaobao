package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallStoreOrderCreateAPIRequest
门店订单创建api API请求
tmall.store.order.create

门店订单创建api */
type TmallStoreOrderCreateAPIRequest struct {
	model.Params
	// 系统自动生成
	_appInfo *AppInfo
	// 创建订单请求
	_createOrderRequest *CreateOrderRequest
}

// New
