package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrechargeaccountflowsget 分页查询储值流水
// alibaba.alsc.crm.recharge.accountflows.get
//
// 增加分页查询储值流水接口
func Alibabaalsccrmrechargeaccountflowsget(clt *core.SDKClient, req *alsc.AlibabaalsccrmrechargeaccountflowsgetAPIRequest, session string) (*alsc.AlibabaalsccrmrechargeaccountflowsgetAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrechargeaccountflowsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
