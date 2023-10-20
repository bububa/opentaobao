package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdeviceopenapigetdevicelist 多条件查询设备分组
// alibaba.campus.device.openapi.getdevicelist
//
// 多条件查询设备分组
func Alibabacampusdeviceopenapigetdevicelist(clt *core.SDKClient, req *campus.AlibabacampusdeviceopenapigetdevicelistAPIRequest, session string) (*campus.AlibabacampusdeviceopenapigetdevicelistAPIResponse, error) {
	var resp campus.AlibabacampusdeviceopenapigetdevicelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
