package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberCheckmerchant 校验商家身份
// alibaba.member.checkmerchant
//
// 校验商家身份
func AlibabaMemberCheckmerchant(clt *core.SDKClient, req *alimember.AlibabaMemberCheckmerchantAPIRequest, resp *alimember.AlibabaMemberCheckmerchantAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
