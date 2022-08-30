package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthSecondardNodeUserSync 二级节点用户数据同步
// alibaba.alihealth.secondard.node.user.sync
//
// 二级节点用户数据同步
func AlibabaAlihealthSecondardNodeUserSync(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthSecondardNodeUserSyncAPIRequest, session string) (*drugtrace.AlibabaAlihealthSecondardNodeUserSyncAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthSecondardNodeUserSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
