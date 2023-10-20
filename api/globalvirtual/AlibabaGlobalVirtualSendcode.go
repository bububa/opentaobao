package globalvirtual

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/globalvirtual"
)

// AlibabaGlobalVirtualSendcode 国际虚拟商品发码服务
// alibaba.global.virtual.sendcode
//
// global virtual send code service
func AlibabaGlobalVirtualSendcode(clt *core.SDKClient, req *globalvirtual.AlibabaGlobalVirtualSendcodeAPIRequest, resp *globalvirtual.AlibabaGlobalVirtualSendcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
