package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusguardcontrollerofflinedata 点位离线数据拉取
// alibaba.campus.guard.controller.offlinedata
//
// 点位离线数据拉取
func Alibabacampusguardcontrollerofflinedata(clt *core.SDKClient, req *campus.AlibabacampusguardcontrollerofflinedataAPIRequest, session string) (*campus.AlibabacampusguardcontrollerofflinedataAPIResponse, error) {
	var resp campus.AlibabacampusguardcontrollerofflinedataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
