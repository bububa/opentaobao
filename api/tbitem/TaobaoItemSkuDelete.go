package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoitemskudelete 删除SKU
// taobao.item.sku.delete
//
// 删除一个sku的数据&lt;br/&gt;需要删除的sku通过属性properties进行匹配查找
func Taobaoitemskudelete(clt *core.SDKClient, req *tbitem.TaobaoitemskudeleteAPIRequest, session string) (*tbitem.TaobaoitemskudeleteAPIResponse, error) {
	var resp tbitem.TaobaoitemskudeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
