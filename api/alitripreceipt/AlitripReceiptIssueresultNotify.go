package alitripreceipt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripreceipt"
)

// AlitripReceiptIssueresultNotify 飞猪发票开票结果通知
// alitrip.receipt.issueresult.notify
//
// 飞猪发票开票结果通知
func AlitripReceiptIssueresultNotify(clt *core.SDKClient, req *alitripreceipt.AlitripReceiptIssueresultNotifyAPIRequest, session string) (*alitripreceipt.AlitripReceiptIssueresultNotifyAPIResponse, error) {
	var resp alitripreceipt.AlitripReceiptIssueresultNotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
