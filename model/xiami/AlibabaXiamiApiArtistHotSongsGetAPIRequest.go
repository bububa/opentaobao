package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiArtistHotSongsGetAPIRequest
热门歌曲 API请求
alibaba.xiami.api.artist.hotSongs.get

热门歌曲 */
type AlibabaXiamiApiArtistHotSongsGetAPIRequest struct {
	model.Params
	// 艺人id
	_id int64
}

// New
