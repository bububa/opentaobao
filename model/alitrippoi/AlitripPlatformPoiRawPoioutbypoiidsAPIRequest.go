package alitrippoi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripPlatformPoiRawPoioutbypoiidsAPIRequest
根据poiId输出飞猪poi数据 API请求
alitrip.platform.poi.raw.poioutbypoiids

根据poiId输出飞猪poi数据 */
type AlitripPlatformPoiRawPoioutbypoiidsAPIRequest struct {
	model.Params
	// 查询参数
	_fliggyPoiidParam *FliggyPoiIdParam
}

// New
