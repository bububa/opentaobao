package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardControllerConfigsync 门禁控制器配置项同步
// alibaba.campus.guard.controller.configsync
//
// 门禁控制器配置项同步
func AlibabaCampusGuardControllerConfigsync(clt *core.SDKClient, req *campus.AlibabaCampusGuardControllerConfigsyncAPIRequest, session string) (*campus.AlibabaCampusGuardControllerConfigsyncAPIResponse, error) {
	var resp campus.AlibabaCampusGuardControllerConfigsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
