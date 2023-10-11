package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaValueCoinIssue 爱豆发放
// alibaba.value.coin.issue
//
// 爱豆发放
func AlibabaValueCoinIssue(clt *core.SDKClient, req *charity.AlibabaValueCoinIssueAPIRequest, session string) (*charity.AlibabaValueCoinIssueAPIResponse, error) {
	var resp charity.AlibabaValueCoinIssueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
