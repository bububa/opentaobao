package xiamicontent

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiamicontent"
)

// XiamiContentSongsAudioGetrefrain 获取副歌信息
// xiami.content.songs.audio.getrefrain
//
// 获取歌曲音频副歌
func XiamiContentSongsAudioGetrefrain(ctx context.Context, clt *core.SDKClient, req *xiamicontent.XiamiContentSongsAudioGetrefrainAPIRequest, resp *xiamicontent.XiamiContentSongsAudioGetrefrainAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
