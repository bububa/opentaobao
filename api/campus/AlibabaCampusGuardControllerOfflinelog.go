package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardControllerOfflinelog 门禁控制器离线日志同步
// alibaba.campus.guard.controller.offlinelog
//
// 门禁控制器离线日志同步
func AlibabaCampusGuardControllerOfflinelog(clt *core.SDKClient, req *campus.AlibabaCampusGuardControllerOfflinelogAPIRequest, session string) (*campus.AlibabaCampusGuardControllerOfflinelogAPIResponse, error) {
	var resp campus.AlibabaCampusGuardControllerOfflinelogAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
