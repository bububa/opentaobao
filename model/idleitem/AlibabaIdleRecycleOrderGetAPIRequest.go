package idleitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerecycleordergetAPIRequest 闲鱼回收订单查询V2 API请求
// alibaba.idle.recycle.order.get
//
// 闲鱼回收业务中,外部回收商作为交易上买家,闲鱼用户下单后,需要回收商主动拉取交易订单
type AlibabaidlerecycleordergetAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// NewAlibabaidlerecycleordergetRequest 初始化AlibabaidlerecycleordergetAPIRequest对象
func NewAlibabaidlerecycleordergetRequest() *AlibabaidlerecycleordergetAPIRequest {
	return &AlibabaidlerecycleordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerecycleordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerecycleordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerecycleordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *AlibabaidlerecycleordergetAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaidlerecycleordergetAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
