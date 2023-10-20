package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabsiotdevicelistget 获取iot设备列表
// alibaba.ailabs.iot.device.list.get
//
// 通过此接口获取用户名下的iot设备列表
func Alibabaailabsiotdevicelistget(clt *core.SDKClient, req *tmallgenie.AlibabaailabsiotdevicelistgetAPIRequest, session string) (*tmallgenie.AlibabaailabsiotdevicelistgetAPIResponse, error) {
	var resp tmallgenie.AlibabaailabsiotdevicelistgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
