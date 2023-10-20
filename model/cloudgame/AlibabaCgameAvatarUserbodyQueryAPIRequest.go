package cloudgame

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCgameAvatarUserbodyQueryAPIRequest 用户Avatar body查询 API请求
// alibaba.cgame.avatar.userbody.query
//
// Avatar用户body数据查询
type AlibabaCgameAvatarUserbodyQueryAPIRequest struct {
	model.Params
	// 查询数据所属用户的mixUserId
	_mixUserId string
}

// NewAlibabaCgameAvatarUserbodyQueryRequest 初始化AlibabaCgameAvatarUserbodyQueryAPIRequest对象
func NewAlibabaCgameAvatarUserbodyQueryRequest() *AlibabaCgameAvatarUserbodyQueryAPIRequest {
	return &AlibabaCgameAvatarUserbodyQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCgameAvatarUserbodyQueryAPIRequest) Reset() {
	r._mixUserId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCgameAvatarUserbodyQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.cgame.avatar.userbody.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCgameAvatarUserbodyQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCgameAvatarUserbodyQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMixUserId is MixUserId Setter
// 查询数据所属用户的mixUserId
func (r *AlibabaCgameAvatarUserbodyQueryAPIRequest) SetMixUserId(_mixUserId string) error {
	r._mixUserId = _mixUserId
	r.Set("mix_user_id", _mixUserId)
	return nil
}

// GetMixUserId MixUserId Getter
func (r AlibabaCgameAvatarUserbodyQueryAPIRequest) GetMixUserId() string {
	return r._mixUserId
}

var poolAlibabaCgameAvatarUserbodyQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCgameAvatarUserbodyQueryRequest()
	},
}

// GetAlibabaCgameAvatarUserbodyQueryRequest 从 sync.Pool 获取 AlibabaCgameAvatarUserbodyQueryAPIRequest
func GetAlibabaCgameAvatarUserbodyQueryAPIRequest() *AlibabaCgameAvatarUserbodyQueryAPIRequest {
	return poolAlibabaCgameAvatarUserbodyQueryAPIRequest.Get().(*AlibabaCgameAvatarUserbodyQueryAPIRequest)
}

// ReleaseAlibabaCgameAvatarUserbodyQueryAPIRequest 将 AlibabaCgameAvatarUserbodyQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaCgameAvatarUserbodyQueryAPIRequest(v *AlibabaCgameAvatarUserbodyQueryAPIRequest) {
	v.Reset()
	poolAlibabaCgameAvatarUserbodyQueryAPIRequest.Put(v)
}
