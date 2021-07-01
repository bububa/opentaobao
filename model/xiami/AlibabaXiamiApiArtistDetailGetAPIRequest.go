package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaXiamiApiArtistDetailGetAPIRequest
艺人详情 API请求
alibaba.xiami.api.artist.detail.get

艺人详情 */
type AlibabaXiamiApiArtistDetailGetAPIRequest struct {
	model.Params
	// 艺人id
	_id int64
	// 是否显示description, show为显示, 其他为不显示
	_description string
}

// New
