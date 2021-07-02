package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaIotDeviceCorpusGet IoT设备支持语料获取
// alibaba.iot.device.corpus.get
//
// ISV通过该接口获取天猫精灵IoT设备支持控制或查询的语料
func AlibabaIotDeviceCorpusGet(clt *core.SDKClient, req *tmallgenie.AlibabaIotDeviceCorpusGetAPIRequest, session string) (*tmallgenie.AlibabaIotDeviceCorpusGetAPIResponse, error) {
	var resp tmallgenie.AlibabaIotDeviceCorpusGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
