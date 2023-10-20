package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusrefundfeegetAPIRequest 查询退票费用明细 API请求
// taobao.bus.refundfee.get
//
// 查询退票的费用信息
type TaobaobusrefundfeegetAPIRequest struct {
	model.Params
	// 飞猪子订单号
	_subOrderIds []int64
	// 飞猪订单号
	_aliTripOrderId string
}

// NewTaobaobusrefundfeegetRequest 初始化TaobaobusrefundfeegetAPIRequest对象
func NewTaobaobusrefundfeegetRequest() *TaobaobusrefundfeegetAPIRequest {
	return &TaobaobusrefundfeegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusrefundfeegetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.refundfee.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusrefundfeegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusrefundfeegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubOrderIds is SubOrderIds Setter
// 飞猪子订单号
func (r *TaobaobusrefundfeegetAPIRequest) SetSubOrderIds(_subOrderIds []int64) error {
	r._subOrderIds = _subOrderIds
	r.Set("sub_order_ids", _subOrderIds)
	return nil
}

// GetSubOrderIds SubOrderIds Getter
func (r TaobaobusrefundfeegetAPIRequest) GetSubOrderIds() []int64 {
	return r._subOrderIds
}

// SetAliTripOrderId is AliTripOrderId Setter
// 飞猪订单号
func (r *TaobaobusrefundfeegetAPIRequest) SetAliTripOrderId(_aliTripOrderId string) error {
	r._aliTripOrderId = _aliTripOrderId
	r.Set("ali_trip_order_id", _aliTripOrderId)
	return nil
}

// GetAliTripOrderId AliTripOrderId Getter
func (r TaobaobusrefundfeegetAPIRequest) GetAliTripOrderId() string {
	return r._aliTripOrderId
}
