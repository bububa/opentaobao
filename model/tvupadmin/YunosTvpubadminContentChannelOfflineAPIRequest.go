package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentchannelofflineAPIRequest 迎客松影视频道下线 API请求
// yunos.tvpubadmin.content.channel.offline
//
// 迎客松影视频道下线
type YunostvpubadmincontentchannelofflineAPIRequest struct {
	model.Params
	// id
	_id int64
}

// NewYunostvpubadmincontentchannelofflineRequest 初始化YunostvpubadmincontentchannelofflineAPIRequest对象
func NewYunostvpubadmincontentchannelofflineRequest() *YunostvpubadmincontentchannelofflineAPIRequest {
	return &YunostvpubadmincontentchannelofflineAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentchannelofflineAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.channel.offline"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentchannelofflineAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentchannelofflineAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// id
func (r *YunostvpubadmincontentchannelofflineAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunostvpubadmincontentchannelofflineAPIRequest) GetId() int64 {
	return r._id
}
