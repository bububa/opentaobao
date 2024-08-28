package category

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// TaobaoItemcatsAuthorizeGet 查询商家被授权品牌列表和类目列表
// taobao.itemcats.authorize.get
//
// 查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表
func TaobaoItemcatsAuthorizeGet(ctx context.Context, clt *core.SDKClient, req *category.TaobaoItemcatsAuthorizeGetAPIRequest, resp *category.TaobaoItemcatsAuthorizeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
