package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

/* AlibabaAlinkDeviceDetailGet
获取设备详情
alibaba.alink.device.detail.get

阿里智能获取设备详情 */
func AlibabaAlinkDeviceDetailGet(clt *core.SDKClient, req *alink.AlibabaAlinkDeviceDetailGetAPIRequest, session string) (*alink.AlibabaAlinkDeviceDetailGetAPIResponse, error) {
	var resp alink.AlibabaAlinkDeviceDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
