package ieagency

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentRefundNewGetlistAPIRequest 新查询退票申请单列表 API请求
// taobao.alitrip.ie.agent.refund.new.getlist
//
// 查询商家国际机票退票申请单列表
type TaobaoAlitripIeAgentRefundNewGetlistAPIRequest struct {
	model.Params
	// 列表查询
	_paramRefundOrderQueryListRq *RefundOrderQueryListRq
}

// NewTaobaoAlitripIeAgentRefundNewGetlistRequest 初始化TaobaoAlitripIeAgentRefundNewGetlistAPIRequest对象
func NewTaobaoAlitripIeAgentRefundNewGetlistRequest() *TaobaoAlitripIeAgentRefundNewGetlistAPIRequest {
	return &TaobaoAlitripIeAgentRefundNewGetlistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripIeAgentRefundNewGetlistAPIRequest) Reset() {
	r._paramRefundOrderQueryListRq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentRefundNewGetlistAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.refund.new.getlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentRefundNewGetlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripIeAgentRefundNewGetlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRefundOrderQueryListRq is ParamRefundOrderQueryListRq Setter
// 列表查询
func (r *TaobaoAlitripIeAgentRefundNewGetlistAPIRequest) SetParamRefundOrderQueryListRq(_paramRefundOrderQueryListRq *RefundOrderQueryListRq) error {
	r._paramRefundOrderQueryListRq = _paramRefundOrderQueryListRq
	r.Set("param_refund_order_query_list_rq", _paramRefundOrderQueryListRq)
	return nil
}

// GetParamRefundOrderQueryListRq ParamRefundOrderQueryListRq Getter
func (r TaobaoAlitripIeAgentRefundNewGetlistAPIRequest) GetParamRefundOrderQueryListRq() *RefundOrderQueryListRq {
	return r._paramRefundOrderQueryListRq
}

var poolTaobaoAlitripIeAgentRefundNewGetlistAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripIeAgentRefundNewGetlistRequest()
	},
}

// GetTaobaoAlitripIeAgentRefundNewGetlistRequest 从 sync.Pool 获取 TaobaoAlitripIeAgentRefundNewGetlistAPIRequest
func GetTaobaoAlitripIeAgentRefundNewGetlistAPIRequest() *TaobaoAlitripIeAgentRefundNewGetlistAPIRequest {
	return poolTaobaoAlitripIeAgentRefundNewGetlistAPIRequest.Get().(*TaobaoAlitripIeAgentRefundNewGetlistAPIRequest)
}

// ReleaseTaobaoAlitripIeAgentRefundNewGetlistAPIRequest 将 TaobaoAlitripIeAgentRefundNewGetlistAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripIeAgentRefundNewGetlistAPIRequest(v *TaobaoAlitripIeAgentRefundNewGetlistAPIRequest) {
	v.Reset()
	poolTaobaoAlitripIeAgentRefundNewGetlistAPIRequest.Put(v)
}
