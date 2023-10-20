package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalsuitjudgementpush 推送裁判登记信息给集团法务系统
// alibaba.legal.suit.judgement.push
//
// isv推送裁判登记信息给集团法务系统
func Alibabalegalsuitjudgementpush(clt *core.SDKClient, req *legalsuit.AlibabalegalsuitjudgementpushAPIRequest, session string) (*legalsuit.AlibabalegalsuitjudgementpushAPIResponse, error) {
	var resp legalsuit.AlibabalegalsuitjudgementpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
