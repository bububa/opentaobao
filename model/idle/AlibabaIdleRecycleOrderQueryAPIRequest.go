package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleOrderQueryAPIRequest 闲鱼回收订单查询V1 API请求
// alibaba.idle.recycle.order.query
//
// 查询回收订单
type AlibabaIdleRecycleOrderQueryAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
	// 手淘商家的淘宝账号id
	_recycleSupplierId int64
}

// NewAlibabaIdleRecycleOrderQueryRequest 初始化AlibabaIdleRecycleOrderQueryAPIRequest对象
func NewAlibabaIdleRecycleOrderQueryRequest() *AlibabaIdleRecycleOrderQueryAPIRequest {
	return &AlibabaIdleRecycleOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRecycleOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// 订单号
func (r *AlibabaIdleRecycleOrderQueryAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaIdleRecycleOrderQueryAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

// SetRecycleSupplierId is RecycleSupplierId Setter
// 手淘商家的淘宝账号id
func (r *AlibabaIdleRecycleOrderQueryAPIRequest) SetRecycleSupplierId(_recycleSupplierId int64) error {
	r._recycleSupplierId = _recycleSupplierId
	r.Set("recycle_supplier_id", _recycleSupplierId)
	return nil
}

// GetRecycleSupplierId RecycleSupplierId Getter
func (r AlibabaIdleRecycleOrderQueryAPIRequest) GetRecycleSupplierId() int64 {
	return r._recycleSupplierId
}
