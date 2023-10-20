package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdeviceopenapigettemplatelist 查询设备模板
// alibaba.campus.device.openapi.gettemplatelist
//
// 查询设备模板信息
func Alibabacampusdeviceopenapigettemplatelist(clt *core.SDKClient, req *campus.AlibabacampusdeviceopenapigettemplatelistAPIRequest, session string) (*campus.AlibabacampusdeviceopenapigettemplatelistAPIResponse, error) {
	var resp campus.AlibabacampusdeviceopenapigettemplatelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
