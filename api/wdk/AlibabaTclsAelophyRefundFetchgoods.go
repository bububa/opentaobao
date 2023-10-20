package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophyrefundfetchgoods saas 售后逆向 商户发起逆向取货
// alibaba.tcls.aelophy.refund.fetchgoods
//
// saas 售后逆向 商户发起逆向取货
func Alibabatclsaelophyrefundfetchgoods(clt *core.SDKClient, req *wdk.AlibabatclsaelophyrefundfetchgoodsAPIRequest, session string) (*wdk.AlibabatclsaelophyrefundfetchgoodsAPIResponse, error) {
	var resp wdk.AlibabatclsaelophyrefundfetchgoodsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
