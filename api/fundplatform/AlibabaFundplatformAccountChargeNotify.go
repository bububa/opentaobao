package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// Alibabafundplatformaccountchargenotify 账户充值成功通知
// alibaba.fundplatform.account.charge.notify
//
// 通知外部业务方充值成功
func Alibabafundplatformaccountchargenotify(clt *core.SDKClient, req *fundplatform.AlibabafundplatformaccountchargenotifyAPIRequest, session string) (*fundplatform.AlibabafundplatformaccountchargenotifyAPIResponse, error) {
	var resp fundplatform.AlibabafundplatformaccountchargenotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
