package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsAligenieOpenvideoScenepush 视频单集场景接入API
// alibaba.ailabs.aligenie.openvideo.scenepush
//
// 视频单集场景接入API
func AlibabaAilabsAligenieOpenvideoScenepush(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieOpenvideoScenepushAPIRequest, resp *tmallgenie.AlibabaAilabsAligenieOpenvideoScenepushAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
