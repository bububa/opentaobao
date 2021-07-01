package alitrippoi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripPlatformPoiRawPoioutAPIRequest
飞猪poi输出 API请求
alitrip.platform.poi.raw.poiout

输出指定城市poi指定信息 */
type AlitripPlatformPoiRawPoioutAPIRequest struct {
	model.Params
	// 查询参数
	_fliggyPoiOutParam *FliggyPoiOutParam
}

// New
