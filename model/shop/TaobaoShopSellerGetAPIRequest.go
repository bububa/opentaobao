package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoShopSellerGetAPIRequest
卖家店铺基础信息查询 API请求
taobao.shop.seller.get

获取卖家店铺的基本信息 */
type TaobaoShopSellerGetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：Shop 结构中的所有字段；多个字段之间用逗号(,)分隔
	_fields string
}

// NewTaobaoShopSellerGetRequest 初始化TaobaoShopSellerGetAPIRequest对象
func NewTaobaoShopSellerGetRequest() *TaobaoShopSellerGetAPIRequest {
	return &TaobaoShopSellerGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoShopSellerGetAPIRequest) GetApiMethodName() string {
	return "taobao.shop.seller.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoShopSellerGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Fields Setter
// 需返回的字段列表。可选值：Shop 结构中的所有字段；多个字段之间用逗号(,)分隔
func (r *TaobaoShopSellerGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// Get Fields Getter
func (r TaobaoShopSellerGetAPIRequest) GetFields() string {
	return r._fields
}
