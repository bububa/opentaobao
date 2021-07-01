package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

/* AlibabaFundplatformAccountChargeNotify
账户充值成功通知
alibaba.fundplatform.account.charge.notify

通知外部业务方充值成功 */
func AlibabaFundplatformAccountChargeNotify(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformAccountChargeNotifyAPIRequest, session string) (*fundplatform.AlibabaFundplatformAccountChargeNotifyAPIResponse, error) {
	var resp fundplatform.AlibabaFundplatformAccountChargeNotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
