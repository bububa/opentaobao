package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemUpdateListingTmall taobao.item.update.listing天猫分流
// taobao.item.update.listing.tmall
//
// * 单个商品上架&lt;br/&gt;* 输入的num_iid必须属于当前会话用户
func TaobaoItemUpdateListingTmall(clt *core.SDKClient, req *tbitem.TaobaoItemUpdateListingTmallAPIRequest, session string) (*tbitem.TaobaoItemUpdateListingTmallAPIResponse, error) {
	var resp tbitem.TaobaoItemUpdateListingTmallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
