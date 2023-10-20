package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusdeviceopenapisaveeventinfoforibos saveeventinfoforibos
// alibaba.campus.device.openapi.saveeventinfoforibos
//
// IBos的事件信息上报与反馈的处理接口
func Alibabacampusdeviceopenapisaveeventinfoforibos(clt *core.SDKClient, req *campus.AlibabacampusdeviceopenapisaveeventinfoforibosAPIRequest, session string) (*campus.AlibabacampusdeviceopenapisaveeventinfoforibosAPIResponse, error) {
	var resp campus.AlibabacampusdeviceopenapisaveeventinfoforibosAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
