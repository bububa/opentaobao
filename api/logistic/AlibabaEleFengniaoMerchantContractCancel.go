package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabaelefengniaomerchantcontractcancel 蜂鸟商户解约接口
// alibaba.ele.fengniao.merchant.contract.cancel
//
// 通过调用此接口，商家及商家下的所有门店解除蜂鸟物流服务
func Alibabaelefengniaomerchantcontractcancel(clt *core.SDKClient, req *logistic.AlibabaelefengniaomerchantcontractcancelAPIRequest, session string) (*logistic.AlibabaelefengniaomerchantcontractcancelAPIResponse, error) {
	var resp logistic.AlibabaelefengniaomerchantcontractcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
