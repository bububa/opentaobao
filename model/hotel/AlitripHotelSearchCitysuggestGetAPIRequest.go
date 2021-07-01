package hotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelSearchCitysuggestGetAPIRequest
城市Suggest接口 API请求
alitrip.hotel.search.citysuggest.get

城市Suggest接口 */
type AlitripHotelSearchCitysuggestGetAPIRequest struct {
	model.Params
	// 用户输入的词
	_keyWords string
	// 用户id
	_userId int64
}

// New
