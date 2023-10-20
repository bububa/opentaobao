package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophyrefundcsapplynew 代客退
// alibaba.tcls.aelophy.refund.csapply.new
//
// 代客退
func Alibabatclsaelophyrefundcsapplynew(clt *core.SDKClient, req *wdk.AlibabatclsaelophyrefundcsapplynewAPIRequest, session string) (*wdk.AlibabatclsaelophyrefundcsapplynewAPIResponse, error) {
	var resp wdk.AlibabatclsaelophyrefundcsapplynewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
