package moscm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/moscm"
)

// AlibabaMosGoodsInventoryGetinventorys 可售库存查询
// alibaba.mos.goods.inventory.getinventorys
//
// 查询商品的可售、在库和占库数量
func AlibabaMosGoodsInventoryGetinventorys(clt *core.SDKClient, req *moscm.AlibabaMosGoodsInventoryGetinventorysAPIRequest, resp *moscm.AlibabaMosGoodsInventoryGetinventorysAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
