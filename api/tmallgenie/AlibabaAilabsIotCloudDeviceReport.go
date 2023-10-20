package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabsiotclouddevicereport 天猫精灵云云接入设备状态、事件上报接口
// alibaba.ailabs.iot.cloud.device.report
//
// 承接天猫精灵云云接入设备的状态、事件上报
func Alibabaailabsiotclouddevicereport(clt *core.SDKClient, req *tmallgenie.AlibabaailabsiotclouddevicereportAPIRequest, session string) (*tmallgenie.AlibabaailabsiotclouddevicereportAPIResponse, error) {
	var resp tmallgenie.AlibabaailabsiotclouddevicereportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
