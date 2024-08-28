package category

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// TaobaoItemcatsGet 获取后台供卖家发布商品的标准商品类目
// taobao.itemcats.get
//
// 获取后台供卖家发布商品的标准商品类目。
func TaobaoItemcatsGet(ctx context.Context, clt *core.SDKClient, req *category.TaobaoItemcatsGetAPIRequest, resp *category.TaobaoItemcatsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
