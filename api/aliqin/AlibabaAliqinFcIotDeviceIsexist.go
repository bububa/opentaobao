package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaaliqinfciotdeviceisexist 判断设备是否存在
// alibaba.aliqin.fc.iot.device.isexist
//
// 判断设备是否存在
func Alibabaaliqinfciotdeviceisexist(clt *core.SDKClient, req *aliqin.AlibabaaliqinfciotdeviceisexistAPIRequest, session string) (*aliqin.AlibabaaliqinfciotdeviceisexistAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfciotdeviceisexistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
