package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleOrderPerformAPIRequest 回收订单履约V2 API请求
// alibaba.idle.recycle.order.perform
//
// 闲鱼回收业务中,外部服务商作为买家 需要驱动交易节点变化
type AlibabaIdleRecycleOrderPerformAPIRequest struct {
	model.Params
	// 参数
	_param0 *RecycleOrderSynDto
}

// NewAlibabaIdleRecycleOrderPerformRequest 初始化AlibabaIdleRecycleOrderPerformAPIRequest对象
func NewAlibabaIdleRecycleOrderPerformRequest() *AlibabaIdleRecycleOrderPerformAPIRequest {
	return &AlibabaIdleRecycleOrderPerformAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRecycleOrderPerformAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleOrderPerformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleOrderPerformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRecycleOrderPerformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *AlibabaIdleRecycleOrderPerformAPIRequest) SetParam0(_param0 *RecycleOrderSynDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaIdleRecycleOrderPerformAPIRequest) GetParam0() *RecycleOrderSynDto {
	return r._param0
}

var poolAlibabaIdleRecycleOrderPerformAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRecycleOrderPerformRequest()
	},
}

// GetAlibabaIdleRecycleOrderPerformRequest 从 sync.Pool 获取 AlibabaIdleRecycleOrderPerformAPIRequest
func GetAlibabaIdleRecycleOrderPerformAPIRequest() *AlibabaIdleRecycleOrderPerformAPIRequest {
	return poolAlibabaIdleRecycleOrderPerformAPIRequest.Get().(*AlibabaIdleRecycleOrderPerformAPIRequest)
}

// ReleaseAlibabaIdleRecycleOrderPerformAPIRequest 将 AlibabaIdleRecycleOrderPerformAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRecycleOrderPerformAPIRequest(v *AlibabaIdleRecycleOrderPerformAPIRequest) {
	v.Reset()
	poolAlibabaIdleRecycleOrderPerformAPIRequest.Put(v)
}
