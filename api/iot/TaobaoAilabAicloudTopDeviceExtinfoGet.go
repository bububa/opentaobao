package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdeviceextinfoget 获取设备扩展信息
// taobao.ailab.aicloud.top.device.extinfo.get
//
// 获取设备扩展信息
func Taobaoailabaicloudtopdeviceextinfoget(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdeviceextinfogetAPIRequest, session string) (*iot.TaobaoailabaicloudtopdeviceextinfogetAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdeviceextinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
