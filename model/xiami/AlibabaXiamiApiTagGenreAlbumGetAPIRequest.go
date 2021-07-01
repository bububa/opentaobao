package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiTagGenreAlbumGetAPIRequest
虾米音乐－风格，流派专辑列表 API请求
alibaba.xiami.api.tag.genre.album.get

虾米音乐－风格，流派专辑列表 */
type AlibabaXiamiApiTagGenreAlbumGetAPIRequest struct {
	model.Params
	// 1:风格，2:流派
	_type int64
	// 风格，流派id
	_id int64
	// 页数
	_page int64
	// 每页数量
	_limit int64
}

// New
