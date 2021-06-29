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
    TaobaoAntifraudRiskuserGetResponse
}

type TaobaoAntifraudRiskuserGetResponse struct {
    XMLName xml.Name `xml:"antifraud_riskuser_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 服务调用成功时, 返回的系统流水号
    
    EventId   string `json:"event_id,omitempty" xml:"event_id,omitempty"`

    
    // 风险分值
    
    Score   string `json:"score,omitempty" xml:"score,omitempty"`

    
    // 风险结果, 为reject, review, pass三者之一
    
    RiskLevel   string `json:"risk_level,omitempty" xml:"risk_level,omitempty"`

    
    // 风险结果详情列表，包含多个风险结果单项
    
    DetailList   []AccountRiskDetail `json:"detail_list,omitempty" xml:"detail_list>account_risk_detail,omitempty"`
    
    
}
