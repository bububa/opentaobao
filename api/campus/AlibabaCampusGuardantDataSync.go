package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardantDataSync 刷卡数据同步
// alibaba.campus.guardant.data.sync
//
// 数据同步门禁系统
func AlibabaCampusGuardantDataSync(clt *core.SDKClient, req *campus.AlibabaCampusGuardantDataSyncAPIRequest, session string) (*campus.AlibabaCampusGuardantDataSyncAPIResponse, error) {
	var resp campus.AlibabaCampusGuardantDataSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
