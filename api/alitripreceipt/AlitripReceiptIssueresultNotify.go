package alitripreceipt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripreceipt"
)

// AlitripReceiptIssueresultNotify 飞猪发票开票结果通知
// alitrip.receipt.issueresult.notify
//
// 飞猪发票开票结果通知
func AlitripReceiptIssueresultNotify(ctx context.Context, clt *core.SDKClient, req *alitripreceipt.AlitripReceiptIssueresultNotifyAPIRequest, resp *alitripreceipt.AlitripReceiptIssueresultNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
