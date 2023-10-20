package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomVtDistributeSendcode 通信业务外放发送验证码
// alibaba.alicom.vt.distribute.sendcode
//
// 通信业务外放发送验证码
func AlibabaAlicomVtDistributeSendcode(clt *core.SDKClient, req *alicom.AlibabaAlicomVtDistributeSendcodeAPIRequest, resp *alicom.AlibabaAlicomVtDistributeSendcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
