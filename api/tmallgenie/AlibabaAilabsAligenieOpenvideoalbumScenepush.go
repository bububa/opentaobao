package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsAligenieOpenvideoalbumScenepush 视频专辑场景接入接口
// alibaba.ailabs.aligenie.openvideoalbum.scenepush
//
// 视频专辑场景接入接口
func AlibabaAilabsAligenieOpenvideoalbumScenepush(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieOpenvideoalbumScenepushAPIRequest, resp *tmallgenie.AlibabaAilabsAligenieOpenvideoalbumScenepushAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
