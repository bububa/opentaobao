package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptraveltradedeliverAPIRequest 飞猪度假-订单发货接口 API请求
// alitrip.travel.trade.deliver
//
// 航旅度假无需物流普通商品发货接口（不支持二次预约商品），只支持子订单级别发货
type AlitriptraveltradedeliverAPIRequest struct {
	model.Params
	// 子订单id
	_subOrderId int64
}

// NewAlitriptraveltradedeliverRequest 初始化AlitriptraveltradedeliverAPIRequest对象
func NewAlitriptraveltradedeliverRequest() *AlitriptraveltradedeliverAPIRequest {
	return &AlitriptraveltradedeliverAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptraveltradedeliverAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.trade.deliver"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptraveltradedeliverAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptraveltradedeliverAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubOrderId is SubOrderId Setter
// 子订单id
func (r *AlitriptraveltradedeliverAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r AlitriptraveltradedeliverAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}
