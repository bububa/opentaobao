package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmmarketingencrypt 加密
// alibaba.alsc.crm.marketing.encrypt
//
// 加密
func Alibabaalsccrmmarketingencrypt(clt *core.SDKClient, req *alsc.AlibabaalsccrmmarketingencryptAPIRequest, session string) (*alsc.AlibabaalsccrmmarketingencryptAPIResponse, error) {
	var resp alsc.AlibabaalsccrmmarketingencryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
