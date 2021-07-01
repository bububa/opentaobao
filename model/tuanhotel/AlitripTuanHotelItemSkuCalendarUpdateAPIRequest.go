package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTuanHotelItemSkuCalendarUpdateAPIRequest
酒店非标套餐商品日历库存宝贝SKU更新接口 API请求
alitrip.tuan.hotel.item.sku.calendar.update

商户对发布的日历库存类型的宝贝套餐价格库存信息进行更新，仅提供日历库存的宝贝SKU的更新功能，skuId须传递商品已存在的skuId，若想进行SKU新增操作，请选择使用alitrip.tuan.hotel.item.sku.update接口。提供增量更新SKU功能，对于日历库存若传递日期信息，参数中若包含某一日期的价格和库存，则对此日期的数据进行覆盖更新，若不传递则保留此日期的价格库存信息。 */
type AlitripTuanHotelItemSkuCalendarUpdateAPIRequest struct {
	model.Params
	// 宝贝ID
	_itemId int64
	// 宝贝所属类目
	_catId int64
	// 暂不支持此接口对SKU的部分属性进行更新，包括以下属性： 套餐名称、价格、原价、库存、间夜、商家编码、人数、使用次数等
	_itemSkuList []TopTuanItemSkuVO
}

// New
