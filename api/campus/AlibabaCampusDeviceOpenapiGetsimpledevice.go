package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdeviceopenapigetsimpledevice 获取单个设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
// alibaba.campus.device.openapi.getsimpledevice
//
// 获取指定设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
func Alibabacampusdeviceopenapigetsimpledevice(clt *core.SDKClient, req *campus.AlibabacampusdeviceopenapigetsimpledeviceAPIRequest, session string) (*campus.AlibabacampusdeviceopenapigetsimpledeviceAPIResponse, error) {
	var resp campus.AlibabacampusdeviceopenapigetsimpledeviceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
