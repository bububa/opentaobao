package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrechargeaccountflowdetailget 储值流水详细
// alibaba.alsc.crm.recharge.account.flowdetail.get
//
// 查询储值流水详细接口
func Alibabaalsccrmrechargeaccountflowdetailget(clt *core.SDKClient, req *alsc.AlibabaalsccrmrechargeaccountflowdetailgetAPIRequest, session string) (*alsc.AlibabaalsccrmrechargeaccountflowdetailgetAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrechargeaccountflowdetailgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
