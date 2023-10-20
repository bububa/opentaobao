package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophybillverificatecallback 翱象ERP核销回调
// alibaba.tcls.aelophy.bill.verificate.callback
//
// 翱象ERP核销回调
func Alibabatclsaelophybillverificatecallback(clt *core.SDKClient, req *wdk.AlibabatclsaelophybillverificatecallbackAPIRequest, session string) (*wdk.AlibabatclsaelophybillverificatecallbackAPIResponse, error) {
	var resp wdk.AlibabatclsaelophybillverificatecallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
