package retail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/retail"
)

// TmallStoreOrderCreate 门店订单创建api
// tmall.store.order.create
//
// 门店订单创建api
func TmallStoreOrderCreate(clt *core.SDKClient, req *retail.TmallStoreOrderCreateAPIRequest, resp *retail.TmallStoreOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
