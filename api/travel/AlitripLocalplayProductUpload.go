package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// AlitripLocalplayProductUpload 当地玩乐 产品维护接口
// alitrip.localplay.product.upload
//
// 当地玩乐（境内当地玩乐/境外玩乐套餐） 产品维护接口。
// 接口同时支持新商品发布 和 现有商品编辑：
// 1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
// 2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
func AlitripLocalplayProductUpload(ctx context.Context, clt *core.SDKClient, req *travel.AlitripLocalplayProductUploadAPIRequest, resp *travel.AlitripLocalplayProductUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
