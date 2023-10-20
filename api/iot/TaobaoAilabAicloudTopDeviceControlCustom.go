package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdevicecontrolcustom 设备控制自定义扩展接口
// taobao.ailab.aicloud.top.device.control.custom
//
// 设备控制自定义扩展接口
func Taobaoailabaicloudtopdevicecontrolcustom(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdevicecontrolcustomAPIRequest, session string) (*iot.TaobaoailabaicloudtopdevicecontrolcustomAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdevicecontrolcustomAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
