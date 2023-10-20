package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoCancelMerchant 商户取消
// alibaba.ele.fengniao.cancel.merchant
//
// 商户取消配送
func AlibabaEleFengniaoCancelMerchant(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoCancelMerchantAPIRequest, resp *logistic.AlibabaEleFengniaoCancelMerchantAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
