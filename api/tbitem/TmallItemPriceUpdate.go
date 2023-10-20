package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallitempriceupdate 天猫商品/SKU价格更新接口
// tmall.item.price.update
//
// 天猫商品/SKU价格更新接口，支持商品、SKU价格同时更新，支持同一商品下的SKU批量更新。
func Tmallitempriceupdate(clt *core.SDKClient, req *tbitem.TmallitempriceupdateAPIRequest, session string) (*tbitem.TmallitempriceupdateAPIResponse, error) {
	var resp tbitem.TmallitempriceupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
