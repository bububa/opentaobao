package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaTclsAelophyMerchantOrderUpload
商家订单数据上传
alibaba.tcls.aelophy.merchant.order.upload

商家订单数据上传 */
func AlibabaTclsAelophyMerchantOrderUpload(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantOrderUploadAPIRequest, session string) (*wdk.AlibabaTclsAelophyMerchantOrderUploadAPIResponse, error) {
	var resp wdk.AlibabaTclsAelophyMerchantOrderUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
