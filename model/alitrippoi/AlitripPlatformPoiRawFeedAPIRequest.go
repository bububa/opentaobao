package alitrippoi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripPlatformPoiRawFeedAPIRequest
存储poi原始数据 API请求
alitrip.platform.poi.raw.feed

对接外部数据源，外部数据推送poi数据到飞猪 */
type AlitripPlatformPoiRawFeedAPIRequest struct {
	model.Params
	// poi存储参数
	_param0 *TripPoiRawSaveParam
}

// New
