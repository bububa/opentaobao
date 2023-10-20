package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentSongsAudioGetrefrain 获取副歌信息
// xiami.content.songs.audio.getrefrain
//
// 获取歌曲音频副歌
func XiamiContentSongsAudioGetrefrain(clt *core.SDKClient, req *xiamicontent.XiamiContentSongsAudioGetrefrainAPIRequest, resp *xiamicontent.XiamiContentSongsAudioGetrefrainAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
