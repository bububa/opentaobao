package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophyrefundcsapplyrender 商家代客售后逆向申请渲染获取
// alibaba.tcls.aelophy.refund.csapplyrender
//
// 提供商家代客售后逆向申请渲染获取的接口
func Alibabatclsaelophyrefundcsapplyrender(clt *core.SDKClient, req *wdk.AlibabatclsaelophyrefundcsapplyrenderAPIRequest, session string) (*wdk.AlibabatclsaelophyrefundcsapplyrenderAPIResponse, error) {
	var resp wdk.AlibabatclsaelophyrefundcsapplyrenderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
