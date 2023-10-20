package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// AlibabaInteractMediaAudio 音频相关鉴权接口
// alibaba.interact.media.audio
//
// 新音频包的鉴权接口
func AlibabaInteractMediaAudio(clt *core.SDKClient, req *mtopopen.AlibabaInteractMediaAudioAPIRequest, resp *mtopopen.AlibabaInteractMediaAudioAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
