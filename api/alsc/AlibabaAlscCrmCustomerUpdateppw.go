package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcustomerupdateppw 修改支付密码
// alibaba.alsc.crm.customer.updateppw
//
// 修改支付密码
func Alibabaalsccrmcustomerupdateppw(clt *core.SDKClient, req *alsc.AlibabaalsccrmcustomerupdateppwAPIRequest, session string) (*alsc.AlibabaalsccrmcustomerupdateppwAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcustomerupdateppwAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
