package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabaelefengniaocancelmerchant 商户取消
// alibaba.ele.fengniao.cancel.merchant
//
// 商户取消配送
func Alibabaelefengniaocancelmerchant(clt *core.SDKClient, req *logistic.AlibabaelefengniaocancelmerchantAPIRequest, session string) (*logistic.AlibabaelefengniaocancelmerchantAPIResponse, error) {
	var resp logistic.AlibabaelefengniaocancelmerchantAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
