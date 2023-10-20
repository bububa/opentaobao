package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerecycleorderperformAPIRequest 回收订单履约V2 API请求
// alibaba.idle.recycle.order.perform
//
// 闲鱼回收业务中,外部服务商作为买家 需要驱动交易节点变化
type AlibabaidlerecycleorderperformAPIRequest struct {
	model.Params
	// 参数
	_param0 *RecycleOrderSynDto
}

// NewAlibabaidlerecycleorderperformRequest 初始化AlibabaidlerecycleorderperformAPIRequest对象
func NewAlibabaidlerecycleorderperformRequest() *AlibabaidlerecycleorderperformAPIRequest {
	return &AlibabaidlerecycleorderperformAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerecycleorderperformAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.recycle.order.perform"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerecycleorderperformAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerecycleorderperformAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *AlibabaidlerecycleorderperformAPIRequest) SetParam0(_param0 *RecycleOrderSynDto) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaidlerecycleorderperformAPIRequest) GetParam0() *RecycleOrderSynDto {
	return r._param0
}
