package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【国际机票】手工预定回填PNR APIRequest
taobao.alitrip.ie.agent.order.hk

代理商通过手工预定PNR，并回填。
*/
type TaobaoAlitripIeAgentOrderHkRequest struct {
    model.Params

    // 代理商ID
    agentId   int64 

    // 回填pnr信息
    writeBackPnrVO   *IeWriteBackPnrVO 

}

func NewTaobaoAlitripIeAgentOrderHkRequest() *TaobaoAlitripIeAgentOrderHkRequest{
    return &TaobaoAlitripIeAgentOrderHkRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripIeAgentOrderHkRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.order.hk"
}

func (r TaobaoAlitripIeAgentOrderHkRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripIeAgentOrderHkRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoAlitripIeAgentOrderHkRequest) GetAgentId() int64 {
    return r.agentId
}

func (r *TaobaoAlitripIeAgentOrderHkRequest) SetWriteBackPnrVO(writeBackPnrVO *IeWriteBackPnrVO) error {
    r.writeBackPnrVO = writeBackPnrVO
    r.Set("write_back_pnr_v_o", writeBackPnrVO)
    return nil
}

func (r TaobaoAlitripIeAgentOrderHkRequest) GetWriteBackPnrVO() *IeWriteBackPnrVO {
    return r.writeBackPnrVO
}

