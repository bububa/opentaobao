package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiArtistAlbumsGetAPIRequest
艺人专辑 API请求
alibaba.xiami.api.artist.albums.get

艺人专辑 */
type AlibabaXiamiApiArtistAlbumsGetAPIRequest struct {
	model.Params
	// 歌曲数量
	_limit int64
	// 页码
	_page int64
	// 艺人id
	_id int64
}

// New
