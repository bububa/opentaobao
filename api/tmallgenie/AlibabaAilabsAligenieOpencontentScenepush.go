package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsAligenieOpencontentScenepush 音频场景接入接口
// alibaba.ailabs.aligenie.opencontent.scenepush
//
// 天猫精灵音频挂靠场景接入
func AlibabaAilabsAligenieOpencontentScenepush(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieOpencontentScenepushAPIRequest, session string) (*tmallgenie.AlibabaAilabsAligenieOpencontentScenepushAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsAligenieOpencontentScenepushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
