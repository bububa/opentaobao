package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemUpdateDelistingTmall taobao.item.update.delisting.tmall
// taobao.item.update.delisting.tmall
//
// * 单个商品下架&lt;br/&gt;    * 输入的num_iid必须属于当前会话用户
func TaobaoItemUpdateDelistingTmall(clt *core.SDKClient, req *tbitem.TaobaoItemUpdateDelistingTmallAPIRequest, resp *tbitem.TaobaoItemUpdateDelistingTmallAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
