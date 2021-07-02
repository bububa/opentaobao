package tvupadmin

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChannelOfflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.channel.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChannelOfflineAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Id Setter
// id
func (r *YunosTvpubadminContentChannelOfflineAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r YunosTvpubadminContentChannelOfflineAPIRequest) GetId() int64 {
	return r._id
}
