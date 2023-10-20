package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTransferpayQuery 闲鱼转账结果查询
// alibaba.idle.transferpay.query
//
// 商家业务 转账支付的结果查询
func AlibabaIdleTransferpayQuery(clt *core.SDKClient, req *idle.AlibabaIdleTransferpayQueryAPIRequest, session string) (*idle.AlibabaIdleTransferpayQueryAPIResponse, error) {
	var resp idle.AlibabaIdleTransferpayQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
