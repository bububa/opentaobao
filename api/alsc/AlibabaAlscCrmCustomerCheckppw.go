package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcustomercheckppw 校验支付密码
// alibaba.alsc.crm.customer.checkppw
//
// 校验支付密码
func Alibabaalsccrmcustomercheckppw(clt *core.SDKClient, req *alsc.AlibabaalsccrmcustomercheckppwAPIRequest, session string) (*alsc.AlibabaalsccrmcustomercheckppwAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcustomercheckppwAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
