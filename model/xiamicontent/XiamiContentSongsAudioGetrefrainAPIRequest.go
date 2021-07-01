package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* XiamiContentSongsAudioGetrefrainAPIRequest
获取副歌信息 API请求
xiami.content.songs.audio.getrefrain

获取歌曲音频副歌 */
type XiamiContentSongsAudioGetrefrainAPIRequest struct {
	model.Params
	// 歌曲ID
	_songIds []int64
}

// New
