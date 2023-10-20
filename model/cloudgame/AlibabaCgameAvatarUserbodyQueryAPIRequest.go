package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacgameavataruserbodyqueryAPIRequest 用户Avatar body查询 API请求
// alibaba.cgame.avatar.userbody.query
//
// Avatar用户body数据查询
type AlibabacgameavataruserbodyqueryAPIRequest struct {
	model.Params
	// 查询数据所属用户的mixUserId
	_mixUserId string
}

// NewAlibabacgameavataruserbodyqueryRequest 初始化AlibabacgameavataruserbodyqueryAPIRequest对象
func NewAlibabacgameavataruserbodyqueryRequest() *AlibabacgameavataruserbodyqueryAPIRequest {
	return &AlibabacgameavataruserbodyqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacgameavataruserbodyqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.avatar.userbody.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacgameavataruserbodyqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacgameavataruserbodyqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMixUserId is MixUserId Setter
// 查询数据所属用户的mixUserId
func (r *AlibabacgameavataruserbodyqueryAPIRequest) SetMixUserId(_mixUserId string) error {
	r._mixUserId = _mixUserId
	r.Set("mix_user_id", _mixUserId)
	return nil
}

// GetMixUserId MixUserId Getter
func (r AlibabacgameavataruserbodyqueryAPIRequest) GetMixUserId() string {
	return r._mixUserId
}
