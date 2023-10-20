package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrechargeunchargecheckget 储值账户退充值校验
// alibaba.alsc.crm.recharge.unchargecheck.get
//
// 储值账户退充值校验接口
func Alibabaalsccrmrechargeunchargecheckget(clt *core.SDKClient, req *alsc.AlibabaalsccrmrechargeunchargecheckgetAPIRequest, session string) (*alsc.AlibabaalsccrmrechargeunchargecheckgetAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrechargeunchargecheckgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
