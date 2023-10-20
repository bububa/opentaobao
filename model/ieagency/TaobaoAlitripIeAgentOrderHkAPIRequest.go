package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripieagentorderhkAPIRequest 【国际机票】手工预定回填PNR API请求
// taobao.alitrip.ie.agent.order.hk
//
// 代理商通过手工预定PNR，并回填。
type TaobaoalitripieagentorderhkAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 回填pnr信息
	_writeBackPnrVO *IeWriteBackPnrVo
}

// NewTaobaoalitripieagentorderhkRequest 初始化TaobaoalitripieagentorderhkAPIRequest对象
func NewTaobaoalitripieagentorderhkRequest() *TaobaoalitripieagentorderhkAPIRequest {
	return &TaobaoalitripieagentorderhkAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripieagentorderhkAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.order.hk"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripieagentorderhkAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripieagentorderhkAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 代理商ID
func (r *TaobaoalitripieagentorderhkAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoalitripieagentorderhkAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetWriteBackPnrVO is WriteBackPnrVO Setter
// 回填pnr信息
func (r *TaobaoalitripieagentorderhkAPIRequest) SetWriteBackPnrVO(_writeBackPnrVO *IeWriteBackPnrVo) error {
	r._writeBackPnrVO = _writeBackPnrVO
	r.Set("write_back_pnr_v_o", _writeBackPnrVO)
	return nil
}

// GetWriteBackPnrVO WriteBackPnrVO Getter
func (r TaobaoalitripieagentorderhkAPIRequest) GetWriteBackPnrVO() *IeWriteBackPnrVo {
	return r._writeBackPnrVO
}
