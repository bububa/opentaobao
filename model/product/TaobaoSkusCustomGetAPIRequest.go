package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSkusCustomGetAPIRequest
根据外部ID取商品SKU API请求
taobao.skus.custom.get

跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku <br/>这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用) */
type TaobaoSkusCustomGetAPIRequest struct {
	model.Params
	// Sku的外部商家ID
	_outerId string
	// 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开
	_fields string
}

// New
