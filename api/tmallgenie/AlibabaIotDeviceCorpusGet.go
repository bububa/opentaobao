package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaiotdevicecorpusget IoT设备支持语料获取
// alibaba.iot.device.corpus.get
//
// ISV通过该接口获取天猫精灵IoT设备支持控制或查询的语料
func Alibabaiotdevicecorpusget(clt *core.SDKClient, req *tmallgenie.AlibabaiotdevicecorpusgetAPIRequest, session string) (*tmallgenie.AlibabaiotdevicecorpusgetAPIResponse, error) {
	var resp tmallgenie.AlibabaiotdevicecorpusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
