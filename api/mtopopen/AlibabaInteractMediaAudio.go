package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// AlibabaInteractMediaAudio 音频相关鉴权接口
// alibaba.interact.media.audio
//
// 新音频包的鉴权接口
func AlibabaInteractMediaAudio(clt *core.SDKClient, req *mtopopen.AlibabaInteractMediaAudioAPIRequest, session string) (*mtopopen.AlibabaInteractMediaAudioAPIResponse, error) {
	var resp mtopopen.AlibabaInteractMediaAudioAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
