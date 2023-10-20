package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdeviceopenapigetdevicerealtimedata 获取指定设备下指定参数的实时值
// alibaba.campus.device.openapi.getdevicerealtimedata
//
// 获取指定设备下指定参数的实时值
func Alibabacampusdeviceopenapigetdevicerealtimedata(clt *core.SDKClient, req *campus.AlibabacampusdeviceopenapigetdevicerealtimedataAPIRequest, session string) (*campus.AlibabacampusdeviceopenapigetdevicerealtimedataAPIResponse, error) {
	var resp campus.AlibabacampusdeviceopenapigetdevicerealtimedataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
