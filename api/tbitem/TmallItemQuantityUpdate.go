package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemQuantityUpdate 天猫商品/SKU库存更新接口
// tmall.item.quantity.update
//
// 天猫商品/SKU库存更新接口；支持商品库存更新；支持同一商品下的SKU批量更新。
func TmallItemQuantityUpdate(clt *core.SDKClient, req *tbitem.TmallItemQuantityUpdateAPIRequest, resp *tbitem.TmallItemQuantityUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
