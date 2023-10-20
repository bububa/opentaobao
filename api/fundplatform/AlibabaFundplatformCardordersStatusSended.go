package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

// AlibabaFundplatformCardordersStatusSended 制卡商通知实体卡发货完成
// alibaba.fundplatform.cardorders.status.sended
//
// 当制卡商将实体卡发货完成后，需要调用该接口，通知我们已发货。
func AlibabaFundplatformCardordersStatusSended(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformCardordersStatusSendedAPIRequest, resp *fundplatform.AlibabaFundplatformCardordersStatusSendedAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
