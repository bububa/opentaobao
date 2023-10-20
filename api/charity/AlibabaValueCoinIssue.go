package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaValueCoinIssue 爱豆发放
// alibaba.value.coin.issue
//
// 爱豆发放
func AlibabaValueCoinIssue(clt *core.SDKClient, req *charity.AlibabaValueCoinIssueAPIRequest, resp *charity.AlibabaValueCoinIssueAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
