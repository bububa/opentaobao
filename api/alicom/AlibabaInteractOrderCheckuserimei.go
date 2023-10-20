package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaInteractOrderCheckuserimei 金融购机验证设备号
// alibaba.interact.order.checkuserimei
//
// 金融购机验证用户身份信息
func AlibabaInteractOrderCheckuserimei(clt *core.SDKClient, req *alicom.AlibabaInteractOrderCheckuserimeiAPIRequest, resp *alicom.AlibabaInteractOrderCheckuserimeiAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
