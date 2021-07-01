package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiRadioSongsGetAPIRequest
获取电台歌曲 API请求
alibaba.xiami.api.radio.songs.get

获取电台songs */
type AlibabaXiamiApiRadioSongsGetAPIRequest struct {
	model.Params
	// 电台类型
	_type int64
	// 电台id
	_oid int64
	// 歌曲个数, 不传为20
	_limit int64
}

// New
