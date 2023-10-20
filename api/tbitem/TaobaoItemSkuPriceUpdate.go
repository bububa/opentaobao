package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoitemskupriceupdate 更新商品SKU的价格
// taobao.item.sku.price.update
//
// 更新商品SKU的价格
func Taobaoitemskupriceupdate(clt *core.SDKClient, req *tbitem.TaobaoitemskupriceupdateAPIRequest, session string) (*tbitem.TaobaoitemskupriceupdateAPIResponse, error) {
	var resp tbitem.TaobaoitemskupriceupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
