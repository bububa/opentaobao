package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusguardcontrollerofflinelog 门禁控制器离线日志同步
// alibaba.campus.guard.controller.offlinelog
//
// 门禁控制器离线日志同步
func Alibabacampusguardcontrollerofflinelog(clt *core.SDKClient, req *campus.AlibabacampusguardcontrollerofflinelogAPIRequest, session string) (*campus.AlibabacampusguardcontrollerofflinelogAPIResponse, error) {
	var resp campus.AlibabacampusguardcontrollerofflinelogAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
