package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// TaobaoVirtualDsfSupplierInterfaceSwitch 虚拟供应商履约接口切换
// taobao.virtual.dsf.supplier.interface.switch
//
// 虚拟供应商履约接口切换
func TaobaoVirtualDsfSupplierInterfaceSwitch(clt *core.SDKClient, req *alicom.TaobaoVirtualDsfSupplierInterfaceSwitchAPIRequest, session string) (*alicom.TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse, error) {
	var resp alicom.TaobaoVirtualDsfSupplierInterfaceSwitchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
