package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRecycleOrderQueryAPIRequest
闲鱼回收订单查询V1 API请求
alibaba.idle.recycle.order.query

查询回收订单 */
type AlibabaIdleRecycleOrderQueryAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
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
func (r AlibabaIdleRecycleOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizOrderId Setter
// 订单号
func (r *AlibabaIdleRecycleOrderQueryAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// Get BizOrderId Getter
func (r AlibabaIdleRecycleOrderQueryAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
