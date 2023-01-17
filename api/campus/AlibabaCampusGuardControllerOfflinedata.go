package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardControllerOfflinedata 点位离线数据拉取
// alibaba.campus.guard.controller.offlinedata
//
// 点位离线数据拉取
func AlibabaCampusGuardControllerOfflinedata(clt *core.SDKClient, req *campus.AlibabaCampusGuardControllerOfflinedataAPIRequest, session string) (*campus.AlibabaCampusGuardControllerOfflinedataAPIResponse, error) {
	var resp campus.AlibabaCampusGuardControllerOfflinedataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
