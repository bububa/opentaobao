package xiamiopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* XiamiApiSongDetailGetAPIRequest
获取歌曲详情 API请求
xiami.api.song.detail.get

获取歌曲详情 */
type XiamiApiSongDetailGetAPIRequest struct {
	model.Params
	// 歌曲id
	_songIds []int64
}

// New
