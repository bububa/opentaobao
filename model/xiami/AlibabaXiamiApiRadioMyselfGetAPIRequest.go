package xiami

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaxiamiapiradiomyselfgetAPIRequest 我的电台 API请求
// alibaba.xiami.api.radio.myself.get
//
// 我的电台
type AlibabaxiamiapiradiomyselfgetAPIRequest struct {
	model.Params
	// 歌曲数量
	_limit int64
}

// NewAlibabaxiamiapiradiomyselfgetRequest 初始化AlibabaxiamiapiradiomyselfgetAPIRequest对象
func NewAlibabaxiamiapiradiomyselfgetRequest() *AlibabaxiamiapiradiomyselfgetAPIRequest {
	return &AlibabaxiamiapiradiomyselfgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaxiamiapiradiomyselfgetAPIRequest) GetApiMethodName() string {
	return "alibaba.xiami.api.radio.myself.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaxiamiapiradiomyselfgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaxiamiapiradiomyselfgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLimit is Limit Setter
// 歌曲数量
func (r *AlibabaxiamiapiradiomyselfgetAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// GetLimit Limit Getter
func (r AlibabaxiamiapiradiomyselfgetAPIRequest) GetLimit() int64 {
	return r._limit
}
