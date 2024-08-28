package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallProductSpecGet 根据产品规格的Id号获取当个的规格信息
// tmall.product.spec.get
//
// 通过当个的spec_id获取到这个产品规格的信息，主要是因为产品规格是要经过审核的，所以通过这个接口可以获取到是否通过审核&lt;br/&gt;通过参看这个ProductSpec的status判断：&lt;br/&gt;1:表示审核通过&lt;br/&gt;3:表示等待审核。&lt;br/&gt;如果你的id找不到数据，那么就是审核被拒绝。
func TmallProductSpecGet(ctx context.Context, clt *core.SDKClient, req *product.TmallProductSpecGetAPIRequest, resp *product.TmallProductSpecGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
