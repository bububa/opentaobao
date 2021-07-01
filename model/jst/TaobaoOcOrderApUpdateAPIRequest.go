package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOcOrderApUpdateAPIRequest
按OC订单分账 API请求
taobao.oc.order.ap.update

对OC订单执行分账操作 */
type TaobaoOcOrderApUpdateAPIRequest struct {
	model.Params
	// 调用创建OC订单接口生成的id
	_ocOrderId int64
}

// NewTaobaoOcOrderApUpdateRequest 初始化TaobaoOcOrderApUpdateAPIRequest对象
func NewTaobaoOcOrderApUpdateRequest() *TaobaoOcOrderApUpdateAPIRequest {
	return &TaobaoOcOrderApUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOcOrderApUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.oc.order.ap.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOcOrderApUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OcOrderId Setter
// 调用创建OC订单接口生成的id
func (r *TaobaoOcOrderApUpdateAPIRequest) SetOcOrderId(_ocOrderId int64) error {
	r._ocOrderId = _ocOrderId
	r.Set("oc_order_id", _ocOrderId)
	return nil
}

// Get OcOrderId Getter
func (r TaobaoOcOrderApUpdateAPIRequest) GetOcOrderId() int64 {
	return r._ocOrderId
}
