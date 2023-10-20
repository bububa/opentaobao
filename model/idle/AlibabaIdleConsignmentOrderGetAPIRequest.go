package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleconsignmentordergetAPIRequest 闲鱼帮卖订单查询 API请求
// alibaba.idle.consignment.order.get
//
// 闲鱼帮卖服务商以闲鱼交易买家身份查询订单信息
type AlibabaidleconsignmentordergetAPIRequest struct {
	model.Params
	// 闲鱼订单ID
	_bizOrderId int64
}

// NewAlibabaidleconsignmentordergetRequest 初始化AlibabaidleconsignmentordergetAPIRequest对象
func NewAlibabaidleconsignmentordergetRequest() *AlibabaidleconsignmentordergetAPIRequest {
	return &AlibabaidleconsignmentordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleconsignmentordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.consignment.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleconsignmentordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleconsignmentordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 闲鱼订单ID
func (r *AlibabaidleconsignmentordergetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaidleconsignmentordergetAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
