package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiArtistMusiclistGetAPIRequest
热门艺人 API请求
alibaba.xiami.api.artist.musiclist.get

热门艺人 */
type AlibabaXiamiApiArtistMusiclistGetAPIRequest struct {
	model.Params
	// 语种, 有all, chinese, musician, english, japanese, korea
	_type string
	// 所有、男、女、组合分别为(all、male、female、combination)
	_order string
}

// New
