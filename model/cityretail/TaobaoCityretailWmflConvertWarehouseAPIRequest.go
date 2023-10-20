package cityretail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocityretailwmflconvertwarehouseAPIRequest 同城零售完美履约转仓 API请求
// taobao.cityretail.wmfl.convert.warehouse
//
// 同城零售完美履约转仓
type TaobaocityretailwmflconvertwarehouseAPIRequest struct {
	model.Params
	// 淘宝交易单id
	_tbOrderId string
}

// NewTaobaocityretailwmflconvertwarehouseRequest 初始化TaobaocityretailwmflconvertwarehouseAPIRequest对象
func NewTaobaocityretailwmflconvertwarehouseRequest() *TaobaocityretailwmflconvertwarehouseAPIRequest {
	return &TaobaocityretailwmflconvertwarehouseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocityretailwmflconvertwarehouseAPIRequest) GetApiMethodName() string {
	return "taobao.cityretail.wmfl.convert.warehouse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocityretailwmflconvertwarehouseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocityretailwmflconvertwarehouseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTbOrderId is TbOrderId Setter
// 淘宝交易单id
func (r *TaobaocityretailwmflconvertwarehouseAPIRequest) SetTbOrderId(_tbOrderId string) error {
	r._tbOrderId = _tbOrderId
	r.Set("tb_order_id", _tbOrderId)
	return nil
}

// GetTbOrderId TbOrderId Getter
func (r TaobaocityretailwmflconvertwarehouseAPIRequest) GetTbOrderId() string {
	return r._tbOrderId
}
