package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* XiamiContentSongsAudioGetAPIRequest
获取歌曲音频 API请求
xiami.content.songs.audio.get

获取歌曲音频 */
type XiamiContentSongsAudioGetAPIRequest struct {
	model.Params
	// 歌曲ID
	_songIds []int64
}

// New
