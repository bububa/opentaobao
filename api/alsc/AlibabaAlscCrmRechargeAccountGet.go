package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrechargeaccountget 查询储值账户信息
// alibaba.alsc.crm.recharge.account.get
//
// 查询储值账户信息接口
func Alibabaalsccrmrechargeaccountget(clt *core.SDKClient, req *alsc.AlibabaalsccrmrechargeaccountgetAPIRequest, session string) (*alsc.AlibabaalsccrmrechargeaccountgetAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrechargeaccountgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
