package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

/* AlibabaLegalSuitJudgementGet
获取裁判登记信息
alibaba.legal.suit.judgement.get

供ISV供应商获取集团法务系统的裁判登记信息 */
func AlibabaLegalSuitJudgementGet(clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitJudgementGetAPIRequest, session string) (*legalsuit.AlibabaLegalSuitJudgementGetAPIResponse, error) {
	var resp legalsuit.AlibabaLegalSuitJudgementGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
