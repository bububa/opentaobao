package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】手工预定回填PNR API请求
taobao.alitrip.ie.agent.order.hk

代理商通过手工预定PNR，并回填。
*/
type TaobaoAlitripIeAgentOrderHkAPIRequest struct {
    model.Params
    // 代理商ID
    _agentId   int64
    // 回填pnr信息
    _writeBackPnrVO   *IeWriteBackPnrVO
}

// 初始化TaobaoAlitripIeAgentOrderHkAPIRequest对象
func NewTaobaoAlitripIeAgentOrderHkRequest() *TaobaoAlitripIeAgentOrderHkAPIRequest{
    return &TaobaoAlitripIeAgentOrderHkAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentOrderHkAPIRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.order.hk"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentOrderHkAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 代理商ID
func (r *TaobaoAlitripIeAgentOrderHkAPIRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoAlitripIeAgentOrderHkAPIRequest) GetAgentId() int64 {
    return r._agentId
}
// WriteBackPnrVO Setter
// 回填pnr信息
func (r *TaobaoAlitripIeAgentOrderHkAPIRequest) SetWriteBackPnrVO(_writeBackPnrVO *IeWriteBackPnrVO) error {
    r._writeBackPnrVO = _writeBackPnrVO
    r.Set("write_back_pnr_v_o", _writeBackPnrVO)
    return nil
}

// WriteBackPnrVO Getter
func (r TaobaoAlitripIeAgentOrderHkAPIRequest) GetWriteBackPnrVO() *IeWriteBackPnrVO {
    return r._writeBackPnrVO
}
