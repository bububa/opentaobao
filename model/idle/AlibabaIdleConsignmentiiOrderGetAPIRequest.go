package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleconsignmentiiordergetAPIRequest 闲鱼寄卖V2订单查询 API请求
// alibaba.idle.consignmentii.order.get
//
// 闲鱼寄卖V2服务商以闲鱼交易买家身份查询订单信息
type AlibabaidleconsignmentiiordergetAPIRequest struct {
	model.Params
	// 闲鱼订单ID
	_bizOrderId int64
}

// NewAlibabaidleconsignmentiiordergetRequest 初始化AlibabaidleconsignmentiiordergetAPIRequest对象
func NewAlibabaidleconsignmentiiordergetRequest() *AlibabaidleconsignmentiiordergetAPIRequest {
	return &AlibabaidleconsignmentiiordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleconsignmentiiordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.consignmentii.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleconsignmentiiordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleconsignmentiiordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 闲鱼订单ID
func (r *AlibabaidleconsignmentiiordergetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaidleconsignmentiiordergetAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
