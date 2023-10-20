package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleappraiseorderqueryAPIRequest 闲鱼验货宝订单详情查询 API请求
// alibaba.idle.appraise.order.query
//
// 鉴定商调用该接口获取订单状态
type AlibabaidleappraiseorderqueryAPIRequest struct {
	model.Params
	// orderId
	_bizOrderId int64
}

// NewAlibabaidleappraiseorderqueryRequest 初始化AlibabaidleappraiseorderqueryAPIRequest对象
func NewAlibabaidleappraiseorderqueryRequest() *AlibabaidleappraiseorderqueryAPIRequest {
	return &AlibabaidleappraiseorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidleappraiseorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.appraise.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidleappraiseorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidleappraiseorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderId is BizOrderId Setter
// orderId
func (r *AlibabaidleappraiseorderqueryAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r AlibabaidleappraiseorderqueryAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
