package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensoraudio 声音
// alibaba.interact.sensor.audio
//
// 客户端声音
func Alibabainteractsensoraudio(clt *core.SDKClient, req *interact.AlibabainteractsensoraudioAPIRequest, session string) (*interact.AlibabainteractsensoraudioAPIResponse, error) {
	var resp interact.AlibabainteractsensoraudioAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
