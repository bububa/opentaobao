package tbtrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest UD效果外投收藏加购明细拉取 API请求
// alibaba.taobao.ud.smart.order.collect.cart.detail.pull
//
// UD效果外投收藏加购明细拉取
type AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest struct {
	model.Params
	// 查询入参
	_orderTopCursorQueryDto *OrderTopCursorQueryDTO2
}

// NewAlibabaTaobaoUdSmartOrderCollectCartDetailPullRequest 初始化AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest对象
func NewAlibabaTaobaoUdSmartOrderCollectCartDetailPullRequest() *AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest {
	return &AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest) Reset() {
	r._orderTopCursorQueryDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest) GetApiMethodName() string {
	return "alibaba.taobao.ud.smart.order.collect.cart.detail.pull"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderTopCursorQueryDto is OrderTopCursorQueryDto Setter
// 查询入参
func (r *AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest) SetOrderTopCursorQueryDto(_orderTopCursorQueryDto *OrderTopCursorQueryDTO2) error {
	r._orderTopCursorQueryDto = _orderTopCursorQueryDto
	r.Set("order_top_cursor_query_dto", _orderTopCursorQueryDto)
	return nil
}

// GetOrderTopCursorQueryDto OrderTopCursorQueryDto Getter
func (r AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest) GetOrderTopCursorQueryDto() *OrderTopCursorQueryDTO2 {
	return r._orderTopCursorQueryDto
}

var poolAlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTaobaoUdSmartOrderCollectCartDetailPullRequest()
	},
}

// GetAlibabaTaobaoUdSmartOrderCollectCartDetailPullRequest 从 sync.Pool 获取 AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest
func GetAlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest() *AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest {
	return poolAlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest.Get().(*AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest)
}

// ReleaseAlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest 将 AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest 放入 sync.Pool
func ReleaseAlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest(v *AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest) {
	v.Reset()
	poolAlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest.Put(v)
}
