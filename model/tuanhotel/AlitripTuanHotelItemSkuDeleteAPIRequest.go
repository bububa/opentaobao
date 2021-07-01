package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTuanHotelItemSkuDeleteAPIRequest
酒店团购套餐商品SKU删除 API请求
alitrip.tuan.hotel.item.sku.delete

商户对发布的宝贝套餐价格库存信息进行删除 */
type AlitripTuanHotelItemSkuDeleteAPIRequest struct {
	model.Params
	// 宝贝ID
	_itemId int64
	// 宝贝所属类目
	_catId int64
	// 要删除的skuId列表
	_itemDeletedSkuIdList []int64
}

// New
