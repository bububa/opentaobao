package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmMarketingEncrypt 加密
// alibaba.alsc.crm.marketing.encrypt
//
// 加密
func AlibabaAlscCrmMarketingEncrypt(clt *core.SDKClient, req *alsc.AlibabaAlscCrmMarketingEncryptAPIRequest, session string) (*alsc.AlibabaAlscCrmMarketingEncryptAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmMarketingEncryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
