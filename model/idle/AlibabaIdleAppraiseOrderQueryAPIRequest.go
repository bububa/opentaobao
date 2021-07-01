package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleAppraiseOrderQueryAPIRequest
闲鱼验货担保订单详情查询V1 API请求
alibaba.idle.appraise.order.query

鉴定商调用该接口获取订单状态 */
type AlibabaIdleAppraiseOrderQueryAPIRequest struct {
	model.Params
	// orderId
	_bizOrderId int64
}

// NewAlibabaIdleAppraiseOrderQueryRequest 初始化AlibabaIdleAppraiseOrderQueryAPIRequest对象
func NewAlibabaIdleAppraiseOrderQueryRequest() *AlibabaIdleAppraiseOrderQueryAPIRequest {
	return &AlibabaIdleAppraiseOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleAppraiseOrderQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.appraise.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleAppraiseOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizOrderId Setter
// orderId
func (r *AlibabaIdleAppraiseOrderQueryAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// Get BizOrderId Getter
func (r AlibabaIdleAppraiseOrderQueryAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
