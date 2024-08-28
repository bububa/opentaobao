package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsAligenieOpenvideoScenepush 视频单集场景接入API
// alibaba.ailabs.aligenie.openvideo.scenepush
//
// 视频单集场景接入API
func AlibabaAilabsAligenieOpenvideoScenepush(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAilabsAligenieOpenvideoScenepushAPIRequest, resp *tmallgenie.AlibabaAilabsAligenieOpenvideoScenepushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
