package travel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// AlitripDaytoursProductUpload 境外一日游/多日游 产品维护接口
// alitrip.daytours.product.upload
//
// 境外一日游/多日游 产品维护接口。
// 接口同时支持新商品发布 和 现有商品编辑：
// 1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
// 2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
func AlitripDaytoursProductUpload(ctx context.Context, clt *core.SDKClient, req *travel.AlitripDaytoursProductUploadAPIRequest, resp *travel.AlitripDaytoursProductUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
