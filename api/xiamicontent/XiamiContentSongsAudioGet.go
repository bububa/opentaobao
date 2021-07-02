package xiamicontent

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentSongsAudioGet 获取歌曲音频
// xiami.content.songs.audio.get
//
// 获取歌曲音频
func XiamiContentSongsAudioGet(clt *core.SDKClient, req *xiamicontent.XiamiContentSongsAudioGetAPIRequest, session string) (*xiamicontent.XiamiContentSongsAudioGetAPIResponse, error) {
	var resp xiamicontent.XiamiContentSongsAudioGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
