package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdevicedetailinfoget 获取设备详细信息
// taobao.ailab.aicloud.top.device.detailinfo.get
//
// 获取设备详细信息
func Taobaoailabaicloudtopdevicedetailinfoget(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdevicedetailinfogetAPIRequest, session string) (*iot.TaobaoailabaicloudtopdevicedetailinfogetAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdevicedetailinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
