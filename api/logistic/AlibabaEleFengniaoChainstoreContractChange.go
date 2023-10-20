package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabaelefengniaochainstorecontractchange 门店改签合同接口
// alibaba.ele.fengniao.chainstore.contract.change
//
// 通过调用接口，门店改签物流服务包
func Alibabaelefengniaochainstorecontractchange(clt *core.SDKClient, req *logistic.AlibabaelefengniaochainstorecontractchangeAPIRequest, session string) (*logistic.AlibabaelefengniaochainstorecontractchangeAPIResponse, error) {
	var resp logistic.AlibabaelefengniaochainstorecontractchangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
