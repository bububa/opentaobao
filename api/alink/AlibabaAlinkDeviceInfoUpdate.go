package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// AlibabaAlinkDeviceInfoUpdate 更新设备昵称等信息
// alibaba.alink.device.info.update
//
// 更新设备昵称等信息
func AlibabaAlinkDeviceInfoUpdate(clt *core.SDKClient, req *alink.AlibabaAlinkDeviceInfoUpdateAPIRequest, session string) (*alink.AlibabaAlinkDeviceInfoUpdateAPIResponse, error) {
	var resp alink.AlibabaAlinkDeviceInfoUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
