package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdeviceopenapigetdevicerealtimelog 根据设备uuid获取设备采集信息
// alibaba.campus.device.openapi.getdevicerealtimelog
//
// 根据设备uuid获取设备采集信息
func Alibabacampusdeviceopenapigetdevicerealtimelog(clt *core.SDKClient, req *campus.AlibabacampusdeviceopenapigetdevicerealtimelogAPIRequest, session string) (*campus.AlibabacampusdeviceopenapigetdevicerealtimelogAPIResponse, error) {
	var resp campus.AlibabacampusdeviceopenapigetdevicerealtimelogAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
