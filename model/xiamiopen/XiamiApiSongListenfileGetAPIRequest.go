package xiamiopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* XiamiApiSongListenfileGetAPIRequest
获取歌曲试听文件 API请求
xiami.api.song.listenfile.get

获取歌曲试听文件 */
type XiamiApiSongListenfileGetAPIRequest struct {
	model.Params
	// 歌曲id
	_songIds []int64
}

// New
