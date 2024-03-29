package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmMarketingEncrypt 加密
// alibaba.alsc.crm.marketing.encrypt
//
// 加密
func AlibabaAlscCrmMarketingEncrypt(clt *core.SDKClient, req *alsc.AlibabaAlscCrmMarketingEncryptAPIRequest, resp *alsc.AlibabaAlscCrmMarketingEncryptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
