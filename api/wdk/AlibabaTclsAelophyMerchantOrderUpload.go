package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophymerchantorderupload 商家订单数据上传
// alibaba.tcls.aelophy.merchant.order.upload
//
// 商家订单数据上传
func Alibabatclsaelophymerchantorderupload(clt *core.SDKClient, req *wdk.AlibabatclsaelophymerchantorderuploadAPIRequest, session string) (*wdk.AlibabatclsaelophymerchantorderuploadAPIResponse, error) {
	var resp wdk.AlibabatclsaelophymerchantorderuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
