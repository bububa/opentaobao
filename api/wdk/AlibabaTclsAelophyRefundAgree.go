package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophyrefundagree saas 售后逆向 商户同意用户逆向申请
// alibaba.tcls.aelophy.refund.agree
//
// saas 售后逆向 商户同意用户逆向申请
func Alibabatclsaelophyrefundagree(clt *core.SDKClient, req *wdk.AlibabatclsaelophyrefundagreeAPIRequest, session string) (*wdk.AlibabatclsaelophyrefundagreeAPIResponse, error) {
	var resp wdk.AlibabatclsaelophyrefundagreeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
