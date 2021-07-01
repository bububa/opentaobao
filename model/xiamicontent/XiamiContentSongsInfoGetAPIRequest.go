package xiamicontent

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* XiamiContentSongsInfoGetAPIRequest
获取歌曲信息 API请求
xiami.content.songs.info.get

(批量)获取歌曲信息 */
type XiamiContentSongsInfoGetAPIRequest struct {
	model.Params
	// 歌曲ID
	_songIds []int64
}

// New
