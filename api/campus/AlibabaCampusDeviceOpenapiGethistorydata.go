package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdeviceopenapigethistorydata 查询设备历史数据
// alibaba.campus.device.openapi.gethistorydata
//
// 查询历史数据的接口
func Alibabacampusdeviceopenapigethistorydata(clt *core.SDKClient, req *campus.AlibabacampusdeviceopenapigethistorydataAPIRequest, session string) (*campus.AlibabacampusdeviceopenapigethistorydataAPIResponse, error) {
	var resp campus.AlibabacampusdeviceopenapigethistorydataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
