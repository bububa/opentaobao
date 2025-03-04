package tbtrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaobaoUdSmartOrderDetailPullAPIRequest UD效果外投订单明细拉取 API请求
// alibaba.taobao.ud.smart.order.detail.pull
//
// UD效果外投订单明细拉取
type AlibabaTaobaoUdSmartOrderDetailPullAPIRequest struct {
	model.Params
	// 查询入参
	_orderTopCursorQueryDto *OrderTopCursorQueryDto
}

// NewAlibabaTaobaoUdSmartOrderDetailPullRequest 初始化AlibabaTaobaoUdSmartOrderDetailPullAPIRequest对象
func NewAlibabaTaobaoUdSmartOrderDetailPullRequest() *AlibabaTaobaoUdSmartOrderDetailPullAPIRequest {
	return &AlibabaTaobaoUdSmartOrderDetailPullAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTaobaoUdSmartOrderDetailPullAPIRequest) Reset() {
	r._orderTopCursorQueryDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTaobaoUdSmartOrderDetailPullAPIRequest) GetApiMethodName() string {
	return "alibaba.taobao.ud.smart.order.detail.pull"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTaobaoUdSmartOrderDetailPullAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTaobaoUdSmartOrderDetailPullAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderTopCursorQueryDto is OrderTopCursorQueryDto Setter
// 查询入参
func (r *AlibabaTaobaoUdSmartOrderDetailPullAPIRequest) SetOrderTopCursorQueryDto(_orderTopCursorQueryDto *OrderTopCursorQueryDto) error {
	r._orderTopCursorQueryDto = _orderTopCursorQueryDto
	r.Set("order_top_cursor_query_dto", _orderTopCursorQueryDto)
	return nil
}

// GetOrderTopCursorQueryDto OrderTopCursorQueryDto Getter
func (r AlibabaTaobaoUdSmartOrderDetailPullAPIRequest) GetOrderTopCursorQueryDto() *OrderTopCursorQueryDto {
	return r._orderTopCursorQueryDto
}

var poolAlibabaTaobaoUdSmartOrderDetailPullAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTaobaoUdSmartOrderDetailPullRequest()
	},
}

// GetAlibabaTaobaoUdSmartOrderDetailPullRequest 从 sync.Pool 获取 AlibabaTaobaoUdSmartOrderDetailPullAPIRequest
func GetAlibabaTaobaoUdSmartOrderDetailPullAPIRequest() *AlibabaTaobaoUdSmartOrderDetailPullAPIRequest {
	return poolAlibabaTaobaoUdSmartOrderDetailPullAPIRequest.Get().(*AlibabaTaobaoUdSmartOrderDetailPullAPIRequest)
}

// ReleaseAlibabaTaobaoUdSmartOrderDetailPullAPIRequest 将 AlibabaTaobaoUdSmartOrderDetailPullAPIRequest 放入 sync.Pool
func ReleaseAlibabaTaobaoUdSmartOrderDetailPullAPIRequest(v *AlibabaTaobaoUdSmartOrderDetailPullAPIRequest) {
	v.Reset()
	poolAlibabaTaobaoUdSmartOrderDetailPullAPIRequest.Put(v)
}
