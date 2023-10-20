package alitripreceipt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripreceipt"
)

// Alitripreceiptissueresultnotify 飞猪发票开票结果通知
// alitrip.receipt.issueresult.notify
//
// 飞猪发票开票结果通知
func Alitripreceiptissueresultnotify(clt *core.SDKClient, req *alitripreceipt.AlitripreceiptissueresultnotifyAPIRequest, session string) (*alitripreceipt.AlitripreceiptissueresultnotifyAPIResponse, error) {
	var resp alitripreceipt.AlitripreceiptissueresultnotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
