package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophybilldetailquery 账单明细接口
// alibaba.tcls.aelophy.bill.detail.query
//
// 账单明细接口
func Alibabatclsaelophybilldetailquery(clt *core.SDKClient, req *wdk.AlibabatclsaelophybilldetailqueryAPIRequest, session string) (*wdk.AlibabatclsaelophybilldetailqueryAPIResponse, error) {
	var resp wdk.AlibabatclsaelophybilldetailqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
