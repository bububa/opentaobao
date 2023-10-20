package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidleagreementpayquery 代扣详情查询
// alibaba.idle.agreement.pay.query
//
// 查询代扣结果
func Alibabaidleagreementpayquery(clt *core.SDKClient, req *idle.AlibabaidleagreementpayqueryAPIRequest, session string) (*idle.AlibabaidleagreementpayqueryAPIResponse, error) {
	var resp idle.AlibabaidleagreementpayqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
