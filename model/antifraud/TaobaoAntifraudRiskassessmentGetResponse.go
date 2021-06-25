package antifraud

import (
    "github.com/bububa/opentaobao/model"
)

/* 
反欺诈风险识别 APIResponse
taobao.antifraud.riskassessment.get

反欺诈服务是阿里大数据风控服务能力的对外输出，通过用户信誉、行为分析精准识别可信用户和风险用户并实时防御，解决交易、支付、活动等关键业务环节存在的欺诈威胁，保护企业品牌和数据，降低企业经济损失
*/
type TaobaoAntifraudRiskassessmentGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAntifraudRiskassessmentGetResponse `json:"taobao_antifraud_riskassessment_get_response,omitempty"`
}

type TaobaoAntifraudRiskassessmentGetResponse struct {

    // result
    RiskResult   *ResultWrapper `json:"risk_result,omitempty"`

}
