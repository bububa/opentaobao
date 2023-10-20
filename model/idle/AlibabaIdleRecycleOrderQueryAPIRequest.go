package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerecycleorderqueryAPIRequest 闲鱼回收订单查询V1 API请求
// alibaba.idle.recycle.order.query
//
// 查询回收订单
type AlibabaidlerecycleorderqueryAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
	// 手淘商家的淘宝账号id
	_recycleSupplierId int64
}

// NewAlibabaidlerecycleorderqueryRequest 初始化AlibabaidlerecycleorderqueryAPIRequest对象
func NewAlibabaidlerecycleorderqueryRequest() *AlibabaidlerecycleorderqueryAPIRequest {
	return &AlibabaidlerecycleorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerecycleorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerecycleorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerecycleorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *AlibabaidlerecycleorderqueryAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaidlerecycleorderqueryAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// SetRecycleSupplierId is RecycleSupplierId Setter
// 手淘商家的淘宝账号id
func (r *AlibabaidlerecycleorderqueryAPIRequest) SetRecycleSupplierId(_recycleSupplierId int64) error {
	r._recycleSupplierId = _recycleSupplierId
	r.Set("recycle_supplier_id", _recycleSupplierId)
	return nil
}

// GetRecycleSupplierId RecycleSupplierId Getter
func (r AlibabaidlerecycleorderqueryAPIRequest) GetRecycleSupplierId() int64 {
	return r._recycleSupplierId
}
