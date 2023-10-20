package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusguardcontrollerconfigsync 门禁控制器配置项同步
// alibaba.campus.guard.controller.configsync
//
// 门禁控制器配置项同步
func Alibabacampusguardcontrollerconfigsync(clt *core.SDKClient, req *campus.AlibabacampusguardcontrollerconfigsyncAPIRequest, session string) (*campus.AlibabacampusguardcontrollerconfigsyncAPIResponse, error) {
	var resp campus.AlibabacampusguardcontrollerconfigsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
