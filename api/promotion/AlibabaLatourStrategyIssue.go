package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaLatourStrategyIssue 阿里巴巴权益发放接口
// alibaba.latour.strategy.issue
//
// 阿里巴巴权益平台权益发放接口
func AlibabaLatourStrategyIssue(clt *core.SDKClient, req *promotion.AlibabaLatourStrategyIssueAPIRequest, resp *promotion.AlibabaLatourStrategyIssueAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
