package tuanhotel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// AlitripTuanHotelItemSkuCalendarUpdate 酒店非标套餐商品日历库存宝贝SKU更新接口
// alitrip.tuan.hotel.item.sku.calendar.update
//
// 商户对发布的日历库存类型的宝贝套餐价格库存信息进行更新，仅提供日历库存的宝贝SKU的更新功能，skuId须传递商品已存在的skuId，若想进行SKU新增操作，请选择使用alitrip.tuan.hotel.item.sku.update接口。提供增量更新SKU功能，对于日历库存若传递日期信息，参数中若包含某一日期的价格和库存，则对此日期的数据进行覆盖更新，若不传递则保留此日期的价格库存信息。
func AlitripTuanHotelItemSkuCalendarUpdate(ctx context.Context, clt *core.SDKClient, req *tuanhotel.AlitripTuanHotelItemSkuCalendarUpdateAPIRequest, resp *tuanhotel.AlitripTuanHotelItemSkuCalendarUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
