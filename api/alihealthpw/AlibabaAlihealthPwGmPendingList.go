package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwGmPendingList 同情用药待审核工单查询接口
// alibaba.alihealth.pw.gm.pending.list
//
// 同情用药待审核工单查询接口，提供给合作方用来查询待处理工单列表
func AlibabaAlihealthPwGmPendingList(clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwGmPendingListAPIRequest, session string) (*alihealthpw.AlibabaAlihealthPwGmPendingListAPIResponse, error) {
	var resp alihealthpw.AlibabaAlihealthPwGmPendingListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
