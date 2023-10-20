package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrechargechargeupdate 储值充值
// alibaba.alsc.crm.recharge.charge.update
//
// 顾客储值账户充值
func Alibabaalsccrmrechargechargeupdate(clt *core.SDKClient, req *alsc.AlibabaalsccrmrechargechargeupdateAPIRequest, session string) (*alsc.AlibabaalsccrmrechargechargeupdateAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrechargechargeupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
