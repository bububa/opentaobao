package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoMerchantContractCancel 蜂鸟商户解约接口
// alibaba.ele.fengniao.merchant.contract.cancel
//
// 通过调用此接口，商家及商家下的所有门店解除蜂鸟物流服务
func AlibabaEleFengniaoMerchantContractCancel(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoMerchantContractCancelAPIRequest, resp *logistic.AlibabaEleFengniaoMerchantContractCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
