package ieagency

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentOrderHkAPIRequest 【国际机票】手工预定回填PNR API请求
// taobao.alitrip.ie.agent.order.hk
//
// 代理商通过手工预定PNR，并回填。
type TaobaoAlitripIeAgentOrderHkAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 回填pnr信息
	_writeBackPnrVO *IeWriteBackPnrVo
}

// NewTaobaoAlitripIeAgentOrderHkRequest 初始化TaobaoAlitripIeAgentOrderHkAPIRequest对象
func NewTaobaoAlitripIeAgentOrderHkRequest() *TaobaoAlitripIeAgentOrderHkAPIRequest {
	return &TaobaoAlitripIeAgentOrderHkAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripIeAgentOrderHkAPIRequest) Reset() {
	r._agentId = 0
	r._writeBackPnrVO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentOrderHkAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.order.hk"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentOrderHkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripIeAgentOrderHkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 代理商ID
func (r *TaobaoAlitripIeAgentOrderHkAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoAlitripIeAgentOrderHkAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetWriteBackPnrVO is WriteBackPnrVO Setter
// 回填pnr信息
func (r *TaobaoAlitripIeAgentOrderHkAPIRequest) SetWriteBackPnrVO(_writeBackPnrVO *IeWriteBackPnrVo) error {
	r._writeBackPnrVO = _writeBackPnrVO
	r.Set("write_back_pnr_v_o", _writeBackPnrVO)
	return nil
}

// GetWriteBackPnrVO WriteBackPnrVO Getter
func (r TaobaoAlitripIeAgentOrderHkAPIRequest) GetWriteBackPnrVO() *IeWriteBackPnrVo {
	return r._writeBackPnrVO
}

var poolTaobaoAlitripIeAgentOrderHkAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripIeAgentOrderHkRequest()
	},
}

// GetTaobaoAlitripIeAgentOrderHkRequest 从 sync.Pool 获取 TaobaoAlitripIeAgentOrderHkAPIRequest
func GetTaobaoAlitripIeAgentOrderHkAPIRequest() *TaobaoAlitripIeAgentOrderHkAPIRequest {
	return poolTaobaoAlitripIeAgentOrderHkAPIRequest.Get().(*TaobaoAlitripIeAgentOrderHkAPIRequest)
}

// ReleaseTaobaoAlitripIeAgentOrderHkAPIRequest 将 TaobaoAlitripIeAgentOrderHkAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripIeAgentOrderHkAPIRequest(v *TaobaoAlitripIeAgentOrderHkAPIRequest) {
	v.Reset()
	poolTaobaoAlitripIeAgentOrderHkAPIRequest.Put(v)
}
