package globalvirtual

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/globalvirtual"
)

// Alibabaglobalvirtualsendcode 国际虚拟商品发码服务
// alibaba.global.virtual.sendcode
//
// global virtual send code service
func Alibabaglobalvirtualsendcode(clt *core.SDKClient, req *globalvirtual.AlibabaglobalvirtualsendcodeAPIRequest, session string) (*globalvirtual.AlibabaglobalvirtualsendcodeAPIResponse, error) {
	var resp globalvirtual.AlibabaglobalvirtualsendcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
