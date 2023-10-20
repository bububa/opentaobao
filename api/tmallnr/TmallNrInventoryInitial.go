package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrInventoryInitial 门店库存初始化前后端商品绑定
// tmall.nr.inventory.initial
//
// 用于门店业务的商品的初始化，前端商品和后端商品绑定，走后端库存模式
func TmallNrInventoryInitial(clt *core.SDKClient, req *tmallnr.TmallNrInventoryInitialAPIRequest, resp *tmallnr.TmallNrInventoryInitialAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
