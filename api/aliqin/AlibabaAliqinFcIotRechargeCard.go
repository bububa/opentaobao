package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// AlibabaaliqinfciotrechargeCard 按终端号订购增值业务
// alibaba.aliqin.fc.iot.rechargeCard
//
// 按终端号订购增值业务
func AlibabaaliqinfciotrechargeCard(clt *core.SDKClient, req *aliqin.AlibabaaliqinfciotrechargeCardAPIRequest, session string) (*aliqin.AlibabaaliqinfciotrechargeCardAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfciotrechargeCardAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
