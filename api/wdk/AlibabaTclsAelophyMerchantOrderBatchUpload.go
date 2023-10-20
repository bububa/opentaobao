package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophymerchantorderbatchupload 商家订单数据批量上传
// alibaba.tcls.aelophy.merchant.order.batch.upload
//
// 商家订单数据上传
func Alibabatclsaelophymerchantorderbatchupload(clt *core.SDKClient, req *wdk.AlibabatclsaelophymerchantorderbatchuploadAPIRequest, session string) (*wdk.AlibabatclsaelophymerchantorderbatchuploadAPIResponse, error) {
	var resp wdk.AlibabatclsaelophymerchantorderbatchuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
