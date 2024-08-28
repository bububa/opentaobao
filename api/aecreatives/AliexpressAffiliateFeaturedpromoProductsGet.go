package aecreatives

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// AliexpressAffiliateFeaturedpromoProductsGet 联盟主题推广活动商品信息获取
// aliexpress.affiliate.featuredpromo.products.get
//
// 根据联盟主题推广活动或主题品库查询对应的商品。如下品库为固定品库，可长期调用。品库类型和名称如下：高佣品（Hot Product）、新品（New Arrival）、热销商品（Best Seller）、每周尖货（weeklydeals）
func AliexpressAffiliateFeaturedpromoProductsGet(ctx context.Context, clt *core.SDKClient, req *aecreatives.AliexpressAffiliateFeaturedpromoProductsGetAPIRequest, resp *aecreatives.AliexpressAffiliateFeaturedpromoProductsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
