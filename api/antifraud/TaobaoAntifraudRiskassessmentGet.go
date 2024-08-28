package antifraud

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/antifraud"
)

// TaobaoAntifraudRiskassessmentGet 反欺诈风险识别
// taobao.antifraud.riskassessment.get
//
// 反欺诈服务是阿里大数据风控服务能力的对外输出，通过用户信誉、行为分析精准识别可信用户和风险用户并实时防御，解决交易、支付、活动等关键业务环节存在的欺诈威胁，保护企业品牌和数据，降低企业经济损失
func TaobaoAntifraudRiskassessmentGet(ctx context.Context, clt *core.SDKClient, req *antifraud.TaobaoAntifraudRiskassessmentGetAPIRequest, resp *antifraud.TaobaoAntifraudRiskassessmentGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
