package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseProjectActivitySync 电商券数据同步
// alibaba.alihouse.project.activity.sync
//
// 电商券数据同步
func AlibabaAlihouseProjectActivitySync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseProjectActivitySyncAPIRequest, session string) (*alihouse.AlibabaAlihouseProjectActivitySyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseProjectActivitySyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
