package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsAligenieOpenvideoScenepush 视频单集场景接入API
// alibaba.ailabs.aligenie.openvideo.scenepush
//
// 视频单集场景接入API
func AlibabaAilabsAligenieOpenvideoScenepush(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieOpenvideoScenepushAPIRequest, session string) (*tmallgenie.AlibabaAilabsAligenieOpenvideoScenepushAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsAligenieOpenvideoScenepushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
