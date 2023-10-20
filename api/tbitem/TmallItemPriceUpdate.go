package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TmallItemPriceUpdate 天猫商品/SKU价格更新接口
// tmall.item.price.update
//
// 天猫商品/SKU价格更新接口，支持商品、SKU价格同时更新，支持同一商品下的SKU批量更新。
func TmallItemPriceUpdate(clt *core.SDKClient, req *tbitem.TmallItemPriceUpdateAPIRequest, session string) (*tbitem.TmallItemPriceUpdateAPIResponse, error) {
	var resp tbitem.TmallItemPriceUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
