package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoChainstoreContractCancel 门店解约接口
// alibaba.ele.fengniao.chainstore.contract.cancel
//
// 调用成功后，门店和蜂鸟解除物流合同，不能再使用此门店推单
func AlibabaEleFengniaoChainstoreContractCancel(ctx context.Context, clt *core.SDKClient, req *logistic.AlibabaEleFengniaoChainstoreContractCancelAPIRequest, resp *logistic.AlibabaEleFengniaoChainstoreContractCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
