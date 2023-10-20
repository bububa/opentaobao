package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// Alibabamembercheckmerchant 校验商家身份
// alibaba.member.checkmerchant
//
// 校验商家身份
func Alibabamembercheckmerchant(clt *core.SDKClient, req *alimember.AlibabamembercheckmerchantAPIRequest, session string) (*alimember.AlibabamembercheckmerchantAPIResponse, error) {
	var resp alimember.AlibabamembercheckmerchantAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
