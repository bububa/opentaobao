package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardControllerConfigsync 门禁控制器配置项同步
// alibaba.campus.guard.controller.configsync
//
// 门禁控制器配置项同步
func AlibabaCampusGuardControllerConfigsync(clt *core.SDKClient, req *campus.AlibabaCampusGuardControllerConfigsyncAPIRequest, resp *campus.AlibabaCampusGuardControllerConfigsyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
