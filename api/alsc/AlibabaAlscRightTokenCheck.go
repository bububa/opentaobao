package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalscrighttokencheck 实物奖品凭证校验
// alibaba.alsc.right.token.check
//
// 实物奖品凭证校验
func Alibabaalscrighttokencheck(clt *core.SDKClient, req *alsc.AlibabaalscrighttokencheckAPIRequest, session string) (*alsc.AlibabaalscrighttokencheckAPIResponse, error) {
	var resp alsc.AlibabaalscrighttokencheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
