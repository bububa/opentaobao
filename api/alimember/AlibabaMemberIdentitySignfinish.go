package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// Alibabamemberidentitysignfinish 签约确认
// alibaba.member.identity.signfinish
//
// 签约确认
func Alibabamemberidentitysignfinish(clt *core.SDKClient, req *alimember.AlibabamemberidentitysignfinishAPIRequest, session string) (*alimember.AlibabamemberidentitysignfinishAPIResponse, error) {
	var resp alimember.AlibabamemberidentitysignfinishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
