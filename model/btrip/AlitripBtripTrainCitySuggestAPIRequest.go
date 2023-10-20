package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripTrainCitySuggestAPIRequest 火车票城市搜索 API请求
// alitrip.btrip.train.city.suggest
//
// 阿里商旅提供火车票搜索接口，方便OA厂商更精准的对接
type AlitripBtripTrainCitySuggestAPIRequest struct {
	model.Params
	// 用户id
	_userId string
	// 搜索关键字
	_keyword string
	// 企业id
	_corpId string
}

// NewAlitripBtripTrainCitySuggestRequest 初始化AlitripBtripTrainCitySuggestAPIRequest对象
func NewAlitripBtripTrainCitySuggestRequest() *AlitripBtripTrainCitySuggestAPIRequest {
	return &AlitripBtripTrainCitySuggestAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripTrainCitySuggestAPIRequest) Reset() {
	r._userId = ""
	r._keyword = ""
	r._corpId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripTrainCitySuggestAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.train.city.suggest"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripTrainCitySuggestAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripTrainCitySuggestAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户id
func (r *AlitripBtripTrainCitySuggestAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlitripBtripTrainCitySuggestAPIRequest) GetUserId() string {
	return r._userId
}

// SetKeyword is Keyword Setter
// 搜索关键字
func (r *AlitripBtripTrainCitySuggestAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r AlitripBtripTrainCitySuggestAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetCorpId is CorpId Setter
// 企业id
func (r *AlitripBtripTrainCitySuggestAPIRequest) SetCorpId(_corpId string) error {
	r._corpId = _corpId
	r.Set("corp_id", _corpId)
	return nil
}

// GetCorpId CorpId Getter
func (r AlitripBtripTrainCitySuggestAPIRequest) GetCorpId() string {
	return r._corpId
}

var poolAlitripBtripTrainCitySuggestAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripTrainCitySuggestRequest()
	},
}

// GetAlitripBtripTrainCitySuggestRequest 从 sync.Pool 获取 AlitripBtripTrainCitySuggestAPIRequest
func GetAlitripBtripTrainCitySuggestAPIRequest() *AlitripBtripTrainCitySuggestAPIRequest {
	return poolAlitripBtripTrainCitySuggestAPIRequest.Get().(*AlitripBtripTrainCitySuggestAPIRequest)
}

// ReleaseAlitripBtripTrainCitySuggestAPIRequest 将 AlitripBtripTrainCitySuggestAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripTrainCitySuggestAPIRequest(v *AlitripBtripTrainCitySuggestAPIRequest) {
	v.Reset()
	poolAlitripBtripTrainCitySuggestAPIRequest.Put(v)
}
