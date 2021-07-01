package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelGereralskuUpdateAPIRequest
发布SKU信息（如果properties重复 则更新） API请求
alitrip.travel.gereralsku.update

发布SKU信息（如果properties重复 则更新） */
type AlitripTravelGereralskuUpdateAPIRequest struct {
	model.Params
	// 淘宝商品ID
	_itemId int64
	// Sku的销售价格，普通商品使用。精确到2位小数;单位:分。如:20007，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
	_price int64
	// Sku的库存数量，普通商品使用。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
	_quantity int64
	// 商家编码
	_outerId string
	// sku销售属性别名；如套餐1 需要调整成其他 需要在这里修改
	_alias []PropertyAliasInfo
	// 商品属性列表；由类目的属性PID和VID组成，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对
	_properties []CatPropInfo
	// SKU的销售价格库存，日历商品使用
	_dateList []DateInventoryAndPrice
}

// New
