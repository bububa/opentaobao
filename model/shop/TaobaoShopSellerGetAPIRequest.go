package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoshopsellergetAPIRequest 卖家店铺基础信息查询 API请求
// taobao.shop.seller.get
//
// 获取卖家店铺的基本信息
type TaobaoshopsellergetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：Shop 结构中的所有字段；多个字段之间用逗号(,)分隔
	_fields string
}

// NewTaobaoshopsellergetRequest 初始化TaobaoshopsellergetAPIRequest对象
func NewTaobaoshopsellergetRequest() *TaobaoshopsellergetAPIRequest {
	return &TaobaoshopsellergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoshopsellergetAPIRequest) GetApiMethodName() string {
	return "taobao.shop.seller.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoshopsellergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoshopsellergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表。可选值：Shop 结构中的所有字段；多个字段之间用逗号(,)分隔
func (r *TaobaoshopsellergetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoshopsellergetAPIRequest) GetFields() string {
	return r._fields
}
