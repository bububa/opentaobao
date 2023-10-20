package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalsuitjudgementget 获取裁判登记信息
// alibaba.legal.suit.judgement.get
//
// 供ISV供应商获取集团法务系统的裁判登记信息
func Alibabalegalsuitjudgementget(clt *core.SDKClient, req *legalsuit.AlibabalegalsuitjudgementgetAPIRequest, session string) (*legalsuit.AlibabalegalsuitjudgementgetAPIResponse, error) {
	var resp legalsuit.AlibabalegalsuitjudgementgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
