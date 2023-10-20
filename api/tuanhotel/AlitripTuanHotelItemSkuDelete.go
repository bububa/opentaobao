package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// AlitripTuanHotelItemSkuDelete 酒店团购套餐商品SKU删除
// alitrip.tuan.hotel.item.sku.delete
//
// 商户对发布的宝贝套餐价格库存信息进行删除
func AlitripTuanHotelItemSkuDelete(clt *core.SDKClient, req *tuanhotel.AlitripTuanHotelItemSkuDeleteAPIRequest, resp *tuanhotel.AlitripTuanHotelItemSkuDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
