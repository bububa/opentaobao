package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcustomerresetppw 重置支付密码
// alibaba.alsc.crm.customer.resetppw
//
// 重置支付密码
func Alibabaalsccrmcustomerresetppw(clt *core.SDKClient, req *alsc.AlibabaalsccrmcustomerresetppwAPIRequest, session string) (*alsc.AlibabaalsccrmcustomerresetppwAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcustomerresetppwAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
