package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniitemCategoryGet 全渠道商品轻发布类目信息
// taobao.omniitem.category.get
//
// 全渠道商品轻发布类目信息
func TaobaoOmniitemCategoryGet(clt *core.SDKClient, req *omniorder.TaobaoOmniitemCategoryGetAPIRequest, resp *omniorder.TaobaoOmniitemCategoryGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
