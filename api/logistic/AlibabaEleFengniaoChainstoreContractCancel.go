package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

/* AlibabaEleFengniaoChainstoreContractCancel
门店解约接口
alibaba.ele.fengniao.chainstore.contract.cancel

调用成功后，门店和蜂鸟解除物流合同，不能再使用此门店推单 */
func AlibabaEleFengniaoChainstoreContractCancel(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoChainstoreContractCancelAPIRequest, session string) (*logistic.AlibabaEleFengniaoChainstoreContractCancelAPIResponse, error) {
	var resp logistic.AlibabaEleFengniaoChainstoreContractCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
