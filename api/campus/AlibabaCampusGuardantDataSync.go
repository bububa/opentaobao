package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardantDataSync 刷卡数据同步
// alibaba.campus.guardant.data.sync
//
// 数据同步门禁系统
func AlibabaCampusGuardantDataSync(clt *core.SDKClient, req *campus.AlibabaCampusGuardantDataSyncAPIRequest, resp *campus.AlibabaCampusGuardantDataSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
