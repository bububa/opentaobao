package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// Youkuottiotstatuspush iot设备状态变化通知接口
// youku.ott.iot.status.push
//
// ott iot设备状态通知
func Youkuottiotstatuspush(clt *core.SDKClient, req *ottpay.YoukuottiotstatuspushAPIRequest, session string) (*ottpay.YoukuottiotstatuspushAPIResponse, error) {
	var resp ottpay.YoukuottiotstatuspushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
