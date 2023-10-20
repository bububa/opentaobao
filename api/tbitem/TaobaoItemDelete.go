package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemDelete 删除单条商品
// taobao.item.delete
//
// 删除单条商品
func TaobaoItemDelete(clt *core.SDKClient, req *tbitem.TaobaoItemDeleteAPIRequest, resp *tbitem.TaobaoItemDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
