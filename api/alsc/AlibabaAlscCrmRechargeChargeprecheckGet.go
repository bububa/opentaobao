package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrechargechargeprecheckget 储值账户充值前校验
// alibaba.alsc.crm.recharge.chargeprecheck.get
//
// 储值账户充值前校验接口
func Alibabaalsccrmrechargechargeprecheckget(clt *core.SDKClient, req *alsc.AlibabaalsccrmrechargechargeprecheckgetAPIRequest, session string) (*alsc.AlibabaalsccrmrechargechargeprecheckgetAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrechargechargeprecheckgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
