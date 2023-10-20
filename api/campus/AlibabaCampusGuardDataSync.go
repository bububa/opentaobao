package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardDataSync 卡巴数据同步
// alibaba.campus.guard.data.sync
//
// 数据同步门禁系统
func AlibabaCampusGuardDataSync(clt *core.SDKClient, req *campus.AlibabaCampusGuardDataSyncAPIRequest, resp *campus.AlibabaCampusGuardDataSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
