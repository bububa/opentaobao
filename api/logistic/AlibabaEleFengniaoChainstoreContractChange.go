package logistic

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// AlibabaEleFengniaoChainstoreContractChange 门店改签合同接口
// alibaba.ele.fengniao.chainstore.contract.change
//
// 通过调用接口，门店改签物流服务包
func AlibabaEleFengniaoChainstoreContractChange(ctx context.Context, clt *core.SDKClient, req *logistic.AlibabaEleFengniaoChainstoreContractChangeAPIRequest, resp *logistic.AlibabaEleFengniaoChainstoreContractChangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
