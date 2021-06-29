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
type TaobaoAlitripIeAgentOrderHkRequest struct {
    model.Params
    // 代理商ID
    _agentId   int64
    // 回填pnr信息
    _writeBackPnrVO   *IeWriteBackPnrVO
}

// 初始化TaobaoAlitripIeAgentOrderHkRequest对象
func NewTaobaoAlitripIeAgentOrderHkRequest() *TaobaoAlitripIeAgentOrderHkRequest{
    return &TaobaoAlitripIeAgentOrderHkRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentOrderHkRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.order.hk"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentOrderHkRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 代理商ID
func (r *TaobaoAlitripIeAgentOrderHkRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoAlitripIeAgentOrderHkRequest) GetAgentId() int64 {
    return r._agentId
}
// WriteBackPnrVO Setter
// 回填pnr信息
func (r *TaobaoAlitripIeAgentOrderHkRequest) SetWriteBackPnrVO(_writeBackPnrVO *IeWriteBackPnrVO) error {
    r._writeBackPnrVO = _writeBackPnrVO
    r.Set("write_back_pnr_v_o", _writeBackPnrVO)
    return nil
}

// WriteBackPnrVO Getter
func (r TaobaoAlitripIeAgentOrderHkRequest) GetWriteBackPnrVO() *IeWriteBackPnrVO {
    return r._writeBackPnrVO
}
