package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemUpdateListing 一口价商品上架
// taobao.item.update.listing
//
// * 单个商品上架&lt;br/&gt;* 输入的num_iid必须属于当前会话用户
func TaobaoItemUpdateListing(clt *core.SDKClient, req *tbitem.TaobaoItemUpdateListingAPIRequest, session string) (*tbitem.TaobaoItemUpdateListingAPIResponse, error) {
	var resp tbitem.TaobaoItemUpdateListingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
