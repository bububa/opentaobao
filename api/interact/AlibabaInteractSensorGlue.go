package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorGlue 视频播放
// alibaba.interact.sensor.glue
//
// 视频播放
func AlibabaInteractSensorGlue(clt *core.SDKClient, req *interact.AlibabaInteractSensorGlueAPIRequest, resp *interact.AlibabaInteractSensorGlueAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
