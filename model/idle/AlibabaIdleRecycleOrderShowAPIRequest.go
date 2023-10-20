package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerecycleordershowAPIRequest 闲鱼回收订单查询V1.1 API请求
// alibaba.idle.recycle.order.show
//
// 查询回收订单
type AlibabaidlerecycleordershowAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// NewAlibabaidlerecycleordershowRequest 初始化AlibabaidlerecycleordershowAPIRequest对象
func NewAlibabaidlerecycleordershowRequest() *AlibabaidlerecycleordershowAPIRequest {
	return &AlibabaidlerecycleordershowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerecycleordershowAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.order.show"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerecycleordershowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerecycleordershowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *AlibabaidlerecycleordershowAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaidlerecycleordershowAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
