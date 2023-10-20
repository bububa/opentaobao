package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrechargeunchargeupdate 充值退款
// alibaba.alsc.crm.recharge.uncharge.update
//
// 充值退款
func Alibabaalsccrmrechargeunchargeupdate(clt *core.SDKClient, req *alsc.AlibabaalsccrmrechargeunchargeupdateAPIRequest, session string) (*alsc.AlibabaalsccrmrechargeunchargeupdateAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrechargeunchargeupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
