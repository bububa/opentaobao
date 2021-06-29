package rail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出票失败关单接口 API请求
alitrip.rail.trade.closeticket

出票成功回调接口
*/
type AlitripRailTradeCloseticketRequest struct {
    model.Params
    // 平台订单号
    tpOrderId   int64
    // 代理商订单号
    agentId   int64
    // 出票失败原因
    errorMsg   string
    // 出票失败码
    errorCode   string
}

// 初始化AlitripRailTradeCloseticketRequest对象
func NewAlitripRailTradeCloseticketRequest() *AlitripRailTradeCloseticketRequest{
    return &AlitripRailTradeCloseticketRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripRailTradeCloseticketRequest) GetApiMethodName() string {
    return "alitrip.rail.trade.closeticket"
}

// IRequest interface 方法, 获取API参数
func (r AlitripRailTradeCloseticketRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TpOrderId Setter
// 平台订单号
func (r *AlitripRailTradeCloseticketRequest) SetTpOrderId(tpOrderId int64) error {
    r.tpOrderId = tpOrderId
    r.Set("tp_order_id", tpOrderId)
    return nil
}

// TpOrderId Getter
func (r AlitripRailTradeCloseticketRequest) GetTpOrderId() int64 {
    return r.tpOrderId
}
// AgentId Setter
// 代理商订单号
func (r *AlitripRailTradeCloseticketRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

// AgentId Getter
func (r AlitripRailTradeCloseticketRequest) GetAgentId() int64 {
    return r.agentId
}
// ErrorMsg Setter
// 出票失败原因
func (r *AlitripRailTradeCloseticketRequest) SetErrorMsg(errorMsg string) error {
    r.errorMsg = errorMsg
    r.Set("error_msg", errorMsg)
    return nil
}

// ErrorMsg Getter
func (r AlitripRailTradeCloseticketRequest) GetErrorMsg() string {
    return r.errorMsg
}
// ErrorCode Setter
// 出票失败码
func (r *AlitripRailTradeCloseticketRequest) SetErrorCode(errorCode string) error {
    r.errorCode = errorCode
    r.Set("error_code", errorCode)
    return nil
}

// ErrorCode Getter
func (r AlitripRailTradeCloseticketRequest) GetErrorCode() string {
    return r.errorCode
}
