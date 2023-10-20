package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightcitysuggestAPIRequest 机票城市搜索 API请求
// alitrip.btrip.flight.city.suggest
//
// 提供机票城市搜索接口，提高OA用户对接效率
type AlitripbtripflightcitysuggestAPIRequest struct {
	model.Params
	// 用户id
	_userId string
	// 搜索关键字
	_keyword string
	// 企业id
	_corpId string
}

// NewAlitripbtripflightcitysuggestRequest 初始化AlitripbtripflightcitysuggestAPIRequest对象
func NewAlitripbtripflightcitysuggestRequest() *AlitripbtripflightcitysuggestAPIRequest {
	return &AlitripbtripflightcitysuggestAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightcitysuggestAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.city.suggest"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightcitysuggestAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightcitysuggestAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户id
func (r *AlitripbtripflightcitysuggestAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlitripbtripflightcitysuggestAPIRequest) GetUserId() string {
	return r._userId
}

// SetKeyword is Keyword Setter
// 搜索关键字
func (r *AlitripbtripflightcitysuggestAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlitripbtripflightcitysuggestAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetCorpId is CorpId Setter
// 企业id
func (r *AlitripbtripflightcitysuggestAPIRequest) SetCorpId(_corpId string) error {
	r._corpId = _corpId
	r.Set("corp_id", _corpId)
	return nil
}

// GetCorpId CorpId Getter
func (r AlitripbtripflightcitysuggestAPIRequest) GetCorpId() string {
	return r._corpId
}
