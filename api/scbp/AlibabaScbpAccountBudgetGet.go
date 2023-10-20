package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpaccountbudgetget 查询日消耗预算
// alibaba.scbp.account.budget.get
//
// 查询日消耗预算
func Alibabascbpaccountbudgetget(clt *core.SDKClient, req *scbp.AlibabascbpaccountbudgetgetAPIRequest, session string) (*scbp.AlibabascbpaccountbudgetgetAPIResponse, error) {
	var resp scbp.AlibabascbpaccountbudgetgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
