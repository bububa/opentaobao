package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseProjectActivitySync 电商券数据同步
// alibaba.alihouse.project.activity.sync
//
// 电商券数据同步
func AlibabaAlihouseProjectActivitySync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseProjectActivitySyncAPIRequest, resp *alihouse.AlibabaAlihouseProjectActivitySyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
