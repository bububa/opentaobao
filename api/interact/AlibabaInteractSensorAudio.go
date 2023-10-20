package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractSensorAudio 声音
// alibaba.interact.sensor.audio
//
// 客户端声音
func AlibabaInteractSensorAudio(clt *core.SDKClient, req *interact.AlibabaInteractSensorAudioAPIRequest, resp *interact.AlibabaInteractSensorAudioAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
