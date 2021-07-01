package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkOrderSyncWithitemAPIRequest
订单和商品同步接口 API请求
alibaba.wdk.order.sync.withitem

轻pos,将订单和商品的信息一起传到盒马这边，进行创单和添加商品处理。 */
type AlibabaWdkOrderSyncWithitemAPIRequest struct {
	model.Params
	// 商家传过来的交易和商品信息
	_posOrderAndItemSync *PosOrderAndItemSyncDo
}

// New
