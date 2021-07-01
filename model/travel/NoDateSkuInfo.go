package travel

// NoDateSkuInfo 结构体
type NoDateSkuInfo struct {
	// sku销售属性别名；如套餐1  需要调整成其他  需要在这里修改
	Alias []PropertyAliasInfo `json:"alias,omitempty" xml:"alias>property_alias_info,omitempty"`
	// sku商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// Sku的销售价格。单位:分。如:20007，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// SKU销售属性列表,若未填写,其他sku信息不会生效；由类目的属性PID和VID组成，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。
	Properties []CatPropInfo `json:"properties,omitempty" xml:"properties>cat_prop_info,omitempty"`
	// Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
