package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripFlightCitySuggestAPIRequest
机票城市搜索 API请求
alitrip.btrip.flight.city.suggest

提供机票城市搜索接口，提高OA用户对接效率 */
type AlitripBtripFlightCitySuggestAPIRequest struct {
	model.Params
	// 用户id
	_userId string
	// 搜索关键字
	_keyword string
	// 企业id
	_corpId string
}

// New
