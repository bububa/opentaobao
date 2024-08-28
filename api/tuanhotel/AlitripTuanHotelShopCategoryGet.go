package tuanhotel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// AlitripTuanHotelShopCategoryGet 商家店铺类目查询
// alitrip.tuan.hotel.shop.category.get
//
// 查询商家店铺类目信息
func AlitripTuanHotelShopCategoryGet(ctx context.Context, clt *core.SDKClient, req *tuanhotel.AlitripTuanHotelShopCategoryGetAPIRequest, resp *tuanhotel.AlitripTuanHotelShopCategoryGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
