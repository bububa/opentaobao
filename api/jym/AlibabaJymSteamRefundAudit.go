package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// Alibabajymsteamrefundaudit 交易猫steam逆向单审核
// alibaba.jym.steam.refund.audit
//
// 交易猫steam逆向单审核
func Alibabajymsteamrefundaudit(clt *core.SDKClient, req *jym.AlibabajymsteamrefundauditAPIRequest, session string) (*jym.AlibabajymsteamrefundauditAPIResponse, error) {
	var resp jym.AlibabajymsteamrefundauditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
