package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// Youkuottiotdevicelistchange iot设备列表变化接口
// youku.ott.iot.devicelist.change
//
// iot设备列表变化接口
func Youkuottiotdevicelistchange(clt *core.SDKClient, req *ottpay.YoukuottiotdevicelistchangeAPIRequest, session string) (*ottpay.YoukuottiotdevicelistchangeAPIResponse, error) {
	var resp ottpay.YoukuottiotdevicelistchangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
