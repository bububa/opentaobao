package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthSecondardNodeUserUpdatestatus 二级节点用户注销
// alibaba.alihealth.secondard.node.user.updatestatus
//
// 注销二级节点用户
func AlibabaAlihealthSecondardNodeUserUpdatestatus(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthSecondardNodeUserUpdatestatusAPIRequest, session string) (*drugtrace.AlibabaAlihealthSecondardNodeUserUpdatestatusAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthSecondardNodeUserUpdatestatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
