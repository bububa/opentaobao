package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemUpdateDelisting 商品下架
// taobao.item.update.delisting
//
// * 单个商品下架&lt;br/&gt;    * 输入的num_iid必须属于当前会话用户
func TaobaoItemUpdateDelisting(clt *core.SDKClient, req *tbitem.TaobaoItemUpdateDelistingAPIRequest, resp *tbitem.TaobaoItemUpdateDelistingAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
