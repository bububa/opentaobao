package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemSkuAdd 添加SKU
// taobao.item.sku.add
//
// 新增一个sku到num_iid指定的商品中 &lt;br/&gt;传入的iid所对应的商品必须属于当前会话的用户
func TaobaoItemSkuAdd(clt *core.SDKClient, req *tbitem.TaobaoItemSkuAddAPIRequest, session string) (*tbitem.TaobaoItemSkuAddAPIResponse, error) {
	var resp tbitem.TaobaoItemSkuAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
