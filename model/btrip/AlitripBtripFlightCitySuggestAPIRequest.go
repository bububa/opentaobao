package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightCitySuggestAPIRequest 机票城市搜索 API请求
// alitrip.btrip.flight.city.suggest
//
// 提供机票城市搜索接口，提高OA用户对接效率
type AlitripBtripFlightCitySuggestAPIRequest struct {
	model.Params
	// 用户id
	_userId string
	// 搜索关键字
	_keyword string
	// 企业id
	_corpId string
}

// NewAlitripBtripFlightCitySuggestRequest 初始化AlitripBtripFlightCitySuggestAPIRequest对象
func NewAlitripBtripFlightCitySuggestRequest() *AlitripBtripFlightCitySuggestAPIRequest {
	return &AlitripBtripFlightCitySuggestAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightCitySuggestAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.city.suggest"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightCitySuggestAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUserId is UserId Setter
// 用户id
func (r *AlitripBtripFlightCitySuggestAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlitripBtripFlightCitySuggestAPIRequest) GetUserId() string {
	return r._userId
}

// SetKeyword is Keyword Setter
// 搜索关键字
func (r *AlitripBtripFlightCitySuggestAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlitripBtripFlightCitySuggestAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetCorpId is CorpId Setter
// 企业id
func (r *AlitripBtripFlightCitySuggestAPIRequest) SetCorpId(_corpId string) error {
	r._corpId = _corpId
	r.Set("corp_id", _corpId)
	return nil
}

// GetCorpId CorpId Getter
func (r AlitripBtripFlightCitySuggestAPIRequest) GetCorpId() string {
	return r._corpId
}
