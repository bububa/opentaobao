package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitJudgementGet 获取裁判登记信息
// alibaba.legal.suit.judgement.get
//
// 供ISV供应商获取集团法务系统的裁判登记信息
func AlibabaLegalSuitJudgementGet(clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitJudgementGetAPIRequest, resp *legalsuit.AlibabaLegalSuitJudgementGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
