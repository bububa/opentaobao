package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChannelOfflineAPIRequest 迎客松影视频道下线 API请求
// yunos.tvpubadmin.content.channel.offline
//
// 迎客松影视频道下线
type YunosTvpubadminContentChannelOfflineAPIRequest struct {
	model.Params
	// id
	_id int64
}

// NewYunosTvpubadminContentChannelOfflineRequest 初始化YunosTvpubadminContentChannelOfflineAPIRequest对象
func NewYunosTvpubadminContentChannelOfflineRequest() *YunosTvpubadminContentChannelOfflineAPIRequest {
	return &YunosTvpubadminContentChannelOfflineAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentChannelOfflineAPIRequest) Reset() {
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChannelOfflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.channel.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChannelOfflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentChannelOfflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// id
func (r *YunosTvpubadminContentChannelOfflineAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunosTvpubadminContentChannelOfflineAPIRequest) GetId() int64 {
	return r._id
}

var poolYunosTvpubadminContentChannelOfflineAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentChannelOfflineRequest()
	},
}

// GetYunosTvpubadminContentChannelOfflineRequest 从 sync.Pool 获取 YunosTvpubadminContentChannelOfflineAPIRequest
func GetYunosTvpubadminContentChannelOfflineAPIRequest() *YunosTvpubadminContentChannelOfflineAPIRequest {
	return poolYunosTvpubadminContentChannelOfflineAPIRequest.Get().(*YunosTvpubadminContentChannelOfflineAPIRequest)
}

// ReleaseYunosTvpubadminContentChannelOfflineAPIRequest 将 YunosTvpubadminContentChannelOfflineAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentChannelOfflineAPIRequest(v *YunosTvpubadminContentChannelOfflineAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentChannelOfflineAPIRequest.Put(v)
}
