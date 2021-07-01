package alitrippoi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripPlatformPoiRawSaverawpoiAPIRequest
POI开放存储能力 API请求
alitrip.platform.poi.raw.saverawpoi

POI开放存储提供离线/在线/纬错更新的能力 */
type AlitripPlatformPoiRawSaverawpoiAPIRequest struct {
	model.Params
	// poi存储参数
	_tripPoiRawSaveParam *TripPoiRawSaveParamV2
}

// New
