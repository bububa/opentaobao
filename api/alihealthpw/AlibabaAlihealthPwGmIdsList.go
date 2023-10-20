package alihealthpw

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwGmIdsList 同情用药根据申请单列表查询申请单
// alibaba.alihealth.pw.gm.ids.list
//
// 同情用药根据申请单列表查询申请单
func AlibabaAlihealthPwGmIdsList(clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwGmIdsListAPIRequest, session string) (*alihealthpw.AlibabaAlihealthPwGmIdsListAPIResponse, error) {
	var resp alihealthpw.AlibabaAlihealthPwGmIdsListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
