package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdeviceopenapigetsimpledevicelist 查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
// alibaba.campus.device.openapi.getsimpledevicelist
//
// 查询设备基础信息集合(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
func Alibabacampusdeviceopenapigetsimpledevicelist(clt *core.SDKClient, req *campus.AlibabacampusdeviceopenapigetsimpledevicelistAPIRequest, session string) (*campus.AlibabacampusdeviceopenapigetsimpledevicelistAPIResponse, error) {
	var resp campus.AlibabacampusdeviceopenapigetsimpledevicelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
