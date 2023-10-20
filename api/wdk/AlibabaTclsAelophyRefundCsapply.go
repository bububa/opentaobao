package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophyrefundcsapply 商家代客售后提交逆向申请
// alibaba.tcls.aelophy.refund.csapply
//
// 商家代客售后提交逆向申请
func Alibabatclsaelophyrefundcsapply(clt *core.SDKClient, req *wdk.AlibabatclsaelophyrefundcsapplyAPIRequest, session string) (*wdk.AlibabatclsaelophyrefundcsapplyAPIResponse, error) {
	var resp wdk.AlibabatclsaelophyrefundcsapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
