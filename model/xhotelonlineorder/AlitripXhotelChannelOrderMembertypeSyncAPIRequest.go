package xhotelonlineorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripXhotelChannelOrderMembertypeSyncAPIRequest 酒店分销渠道会员类型同步 API请求
// alitrip.xhotel.channel.order.membertype.sync
//
// 酒店分销渠道会员类型同步
type AlitripXhotelChannelOrderMembertypeSyncAPIRequest struct {
	model.Params
	// 入参
	_channelSyncOrderMemberType *ChannelSyncOrderMemberType
}

// NewAlitripXhotelChannelOrderMembertypeSyncRequest 初始化AlitripXhotelChannelOrderMembertypeSyncAPIRequest对象
func NewAlitripXhotelChannelOrderMembertypeSyncRequest() *AlitripXhotelChannelOrderMembertypeSyncAPIRequest {
	return &AlitripXhotelChannelOrderMembertypeSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripXhotelChannelOrderMembertypeSyncAPIRequest) Reset() {
	r._channelSyncOrderMemberType = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripXhotelChannelOrderMembertypeSyncAPIRequest) GetApiMethodName() string {
	return "alitrip.xhotel.channel.order.membertype.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripXhotelChannelOrderMembertypeSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripXhotelChannelOrderMembertypeSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelSyncOrderMemberType is ChannelSyncOrderMemberType Setter
// 入参
func (r *AlitripXhotelChannelOrderMembertypeSyncAPIRequest) SetChannelSyncOrderMemberType(_channelSyncOrderMemberType *ChannelSyncOrderMemberType) error {
	r._channelSyncOrderMemberType = _channelSyncOrderMemberType
	r.Set("channel_sync_order_member_type", _channelSyncOrderMemberType)
	return nil
}

// GetChannelSyncOrderMemberType ChannelSyncOrderMemberType Getter
func (r AlitripXhotelChannelOrderMembertypeSyncAPIRequest) GetChannelSyncOrderMemberType() *ChannelSyncOrderMemberType {
	return r._channelSyncOrderMemberType
}

var poolAlitripXhotelChannelOrderMembertypeSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripXhotelChannelOrderMembertypeSyncRequest()
	},
}

// GetAlitripXhotelChannelOrderMembertypeSyncRequest 从 sync.Pool 获取 AlitripXhotelChannelOrderMembertypeSyncAPIRequest
func GetAlitripXhotelChannelOrderMembertypeSyncAPIRequest() *AlitripXhotelChannelOrderMembertypeSyncAPIRequest {
	return poolAlitripXhotelChannelOrderMembertypeSyncAPIRequest.Get().(*AlitripXhotelChannelOrderMembertypeSyncAPIRequest)
}

// ReleaseAlitripXhotelChannelOrderMembertypeSyncAPIRequest 将 AlitripXhotelChannelOrderMembertypeSyncAPIRequest 放入 sync.Pool
func ReleaseAlitripXhotelChannelOrderMembertypeSyncAPIRequest(v *AlitripXhotelChannelOrderMembertypeSyncAPIRequest) {
	v.Reset()
	poolAlitripXhotelChannelOrderMembertypeSyncAPIRequest.Put(v)
}
