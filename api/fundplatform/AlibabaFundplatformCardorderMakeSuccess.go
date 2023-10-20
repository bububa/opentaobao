package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformCardorderMakeSuccess 通知制卡成功
// alibaba.fundplatform.cardorder.make.success
//
// 当外部业务方调用资金平台异步制卡接口后，资金平台制卡成功后通知异步通知业务方
func AlibabaFundplatformCardorderMakeSuccess(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardorderMakeSuccessAPIRequest, resp *fundplatform.AlibabaFundplatformCardorderMakeSuccessAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
