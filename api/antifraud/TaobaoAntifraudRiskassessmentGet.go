package antifraud

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/antifraud"
)

/* 
反欺诈风险识别 
taobao.antifraud.riskassessment.get

反欺诈服务是阿里大数据风控服务能力的对外输出，通过用户信誉、行为分析精准识别可信用户和风险用户并实时防御，解决交易、支付、活动等关键业务环节存在的欺诈威胁，保护企业品牌和数据，降低企业经济损失
*/
func TaobaoAntifraudRiskassessmentGet(clt *core.SDKClient, req *antifraud.TaobaoAntifraudRiskassessmentGetRequest, session string) (*antifraud.TaobaoAntifraudRiskassessmentGetAPIResponse, error) {
    var resp antifraud.TaobaoAntifraudRiskassessmentGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
