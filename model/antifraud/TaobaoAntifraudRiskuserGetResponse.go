package antifraud

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
反欺诈用户风险查询 APIResponse
taobao.antifraud.riskuser.get

根据用户基础信息，核实平台上的用户是否存在欺诈风险
*/
type TaobaoAntifraudRiskuserGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"antifraud_riskuser_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 服务调用成功时, 返回的系统流水号
    
    EventId   string `json:"event_id,omitempty" xml:"