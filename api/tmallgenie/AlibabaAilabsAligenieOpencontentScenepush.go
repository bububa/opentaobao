package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsAligenieOpencontentScenepush 音频场景接入接口
// alibaba.ailabs.aligenie.opencontent.scenepush
//
// 天猫精灵音频挂靠场景接入
func AlibabaAilabsAligenieOpencontentScenepush(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieOpencontentScenepushAPIRequest, resp *tmallgenie.AlibabaAilabsAligenieOpencontentScenepushAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
