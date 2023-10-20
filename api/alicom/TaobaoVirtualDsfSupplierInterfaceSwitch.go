package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Taobaovirtualdsfsupplierinterfaceswitch 虚拟供应商履约接口切换
// taobao.virtual.dsf.supplier.interface.switch
//
// 虚拟供应商履约接口切换
func Taobaovirtualdsfsupplierinterfaceswitch(clt *core.SDKClient, req *alicom.TaobaovirtualdsfsupplierinterfaceswitchAPIRequest, session string) (*alicom.TaobaovirtualdsfsupplierinterfaceswitchAPIResponse, error) {
	var resp alicom.TaobaovirtualdsfsupplierinterfaceswitchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
