package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophyrefunddisagree saas 售后逆向 商户拒绝用户逆向申请
// alibaba.tcls.aelophy.refund.disagree
//
// saas 售后逆向 商户拒绝用户逆向申请
func Alibabatclsaelophyrefunddisagree(clt *core.SDKClient, req *wdk.AlibabatclsaelophyrefunddisagreeAPIRequest, session string) (*wdk.AlibabatclsaelophyrefunddisagreeAPIResponse, error) {
	var resp wdk.AlibabatclsaelophyrefunddisagreeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
