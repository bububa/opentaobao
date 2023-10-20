package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoitemskuadd 添加SKU
// taobao.item.sku.add
//
// 新增一个sku到num_iid指定的商品中 &lt;br/&gt;传入的iid所对应的商品必须属于当前会话的用户
func Taobaoitemskuadd(clt *core.SDKClient, req *tbitem.TaobaoitemskuaddAPIRequest, session string) (*tbitem.TaobaoitemskuaddAPIResponse, error) {
	var resp tbitem.TaobaoitemskuaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
