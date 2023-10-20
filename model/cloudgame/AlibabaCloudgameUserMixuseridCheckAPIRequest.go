package cloudgame

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCloudgameUserMixuseridCheckAPIRequest 云游戏混淆用户ID校验 API请求
// alibaba.cloudgame.user.mixuserid.check
//
// 验证混淆用户ID是否合法
type AlibabaCloudgameUserMixuseridCheckAPIRequest struct {
	model.Params
	// 云游戏混淆用户ID
	_mixUserId string
}

// NewAlibabaCloudgameUserMixuseridCheckRequest 初始化AlibabaCloudgameUserMixuseridCheckAPIRequest对象
func NewAlibabaCloudgameUserMixuseridCheckRequest() *AlibabaCloudgameUserMixuseridCheckAPIRequest {
	return &AlibabaCloudgameUserMixuseridCheckAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCloudgameUserMixuseridCheckAPIRequest) Reset() {
	r._mixUserId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameUserMixuseridCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.user.mixuserid.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameUserMixuseridCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCloudgameUserMixuseridCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMixUserId is MixUserId Setter
// 云游戏混淆用户ID
func (r *AlibabaCloudgameUserMixuseridCheckAPIRequest) SetMixUserId(_mixUserId string) error {
	r._mixUserId = _mixUserId
	r.Set("mix_user_id", _mixUserId)
	return nil
}

// GetMixUserId MixUserId Getter
func (r AlibabaCloudgameUserMixuseridCheckAPIRequest) GetMixUserId() string {
	return r._mixUserId
}

var poolAlibabaCloudgameUserMixuseridCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCloudgameUserMixuseridCheckRequest()
	},
}

// GetAlibabaCloudgameUserMixuseridCheckRequest 从 sync.Pool 获取 AlibabaCloudgameUserMixuseridCheckAPIRequest
func GetAlibabaCloudgameUserMixuseridCheckAPIRequest() *AlibabaCloudgameUserMixuseridCheckAPIRequest {
	return poolAlibabaCloudgameUserMixuseridCheckAPIRequest.Get().(*AlibabaCloudgameUserMixuseridCheckAPIRequest)
}

// ReleaseAlibabaCloudgameUserMixuseridCheckAPIRequest 将 AlibabaCloudgameUserMixuseridCheckAPIRequest 放入 sync.Pool
func ReleaseAlibabaCloudgameUserMixuseridCheckAPIRequest(v *AlibabaCloudgameUserMixuseridCheckAPIRequest) {
	v.Reset()
	poolAlibabaCloudgameUserMixuseridCheckAPIRequest.Put(v)
}
