package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabaelefengniaochainstorecontractcancel 门店解约接口
// alibaba.ele.fengniao.chainstore.contract.cancel
//
// 调用成功后，门店和蜂鸟解除物流合同，不能再使用此门店推单
func Alibabaelefengniaochainstorecontractcancel(clt *core.SDKClient, req *logistic.AlibabaelefengniaochainstorecontractcancelAPIRequest, session string) (*logistic.AlibabaelefengniaochainstorecontractcancelAPIResponse, error) {
	var resp logistic.AlibabaelefengniaochainstorecontractcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
