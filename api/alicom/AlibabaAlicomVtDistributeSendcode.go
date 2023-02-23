package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAlicomVtDistributeSendcode 通信业务外放发送验证码
// alibaba.alicom.vt.distribute.sendcode
//
// 通信业务外放发送验证码
func AlibabaAlicomVtDistributeSendcode(clt *core.SDKClient, req *alicom.AlibabaAlicomVtDistributeSendcodeAPIRequest, session string) (*alicom.AlibabaAlicomVtDistributeSendcodeAPIResponse, error) {
	var resp alicom.AlibabaAlicomVtDistributeSendcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
