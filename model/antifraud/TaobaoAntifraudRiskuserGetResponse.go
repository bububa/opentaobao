package antifraud

import (
    "github.com/bububa/opentaobao/model"
)

/* 
反欺诈用户风险查询 APIResponse
taobao.antifraud.riskuser.get

根据用户基础信息，核实平台上的用户是否存在欺诈风险
*/
type TaobaoAntifraudRiskuserGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAntifraudRiskuserGetResponse `json:"antifraud_riskuser_get_response,omitempty"` 
    TaobaoAntifraudRiskuserGetResponse
}

/* model for simplify = false
type TaobaoAntifraudRiskuserGetResponse struct {

    // 服务调用成功时, 返回的系统流水号
    
    EventId   string `json:"event_id,omitempty"`
    

    // 风险分值
    
    Score   string `json:"score,omitempty"`
    

    // 风险结果, 为reject, review, pass三者之一
    
    RiskLevel   string `json:"risk_level,omitempty"`
    

    // 风险结果详情列表，包含多个风险结果单项
    
    DetailList  struct {
        AccountRiskDetail  []AccountRiskDetail `json:"account_risk_detail,omitempty"`
    } `json:"detail_list,omitempty"`
    

}
*/

type TaobaoAntifraudRiskuserGetResponse struct {

    // 服务调用成功时, 返回的系统流水号
    EventId   string `json:"event_id,omitempty"`

    // 风险分值
    Score   string `json:"score,omitempty"`

    // 风险结果, 为reject, review, pass三者之一
    RiskLevel   string `json:"risk_level,omitempty"`

    // 风险结果详情列表，包含多个风险结果单项
    DetailList   []AccountRiskDetail `json:"detail_list,omitempty"`

}
