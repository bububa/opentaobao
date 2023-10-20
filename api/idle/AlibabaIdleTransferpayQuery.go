package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidletransferpayquery 闲鱼转账结果查询
// alibaba.idle.transferpay.query
//
// 商家业务 转账支付的结果查询
func Alibabaidletransferpayquery(clt *core.SDKClient, req *idle.AlibabaidletransferpayqueryAPIRequest, session string) (*idle.AlibabaidletransferpayqueryAPIResponse, error) {
	var resp idle.AlibabaidletransferpayqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
