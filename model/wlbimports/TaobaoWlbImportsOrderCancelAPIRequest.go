package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbImportsOrderCancelAPIRequest
一般进口取消物流订单 API请求
taobao.wlb.imports.order.cancel

商家在发货后，需要对订单进行取消，如果仓库已经收货则无法取消。 */
type TaobaoWlbImportsOrderCancelAPIRequest struct {
	model.Params
	// 物流订单编号
	_lgorderCode string
}

// New
