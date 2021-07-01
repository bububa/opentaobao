package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiTagGenreSongGetAPIRequest
虾米-风格，流派歌曲 API请求
alibaba.xiami.api.tag.genre.song.get

虾米-风格，流派歌曲 */
type AlibabaXiamiApiTagGenreSongGetAPIRequest struct {
	model.Params
	// 1:风格，2：流派
	_type string
	// 风格或流派id
	_id int64
	// 页数
	_page int64
	// 每页数量
	_limit int64
}

// New
