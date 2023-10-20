package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripxhotelchannelordermembertypesyncAPIRequest 酒店分销渠道会员类型同步 API请求
// alitrip.xhotel.channel.order.membertype.sync
//
// 酒店分销渠道会员类型同步
type AlitripxhotelchannelordermembertypesyncAPIRequest struct {
	model.Params
	// 入参
	_channelSyncOrderMemberType *ChannelSyncOrderMemberType
}

// NewAlitripxhotelchannelordermembertypesyncRequest 初始化AlitripxhotelchannelordermembertypesyncAPIRequest对象
func NewAlitripxhotelchannelordermembertypesyncRequest() *AlitripxhotelchannelordermembertypesyncAPIRequest {
	return &AlitripxhotelchannelordermembertypesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripxhotelchannelordermembertypesyncAPIRequest) GetApiMethodName() string {
	return "alitrip.xhotel.channel.order.membertype.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripxhotelchannelordermembertypesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripxhotelchannelordermembertypesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelSyncOrderMemberType is ChannelSyncOrderMemberType Setter
// 入参
func (r *AlitripxhotelchannelordermembertypesyncAPIRequest) SetChannelSyncOrderMemberType(_channelSyncOrderMemberType *ChannelSyncOrderMemberType) error {
	r._channelSyncOrderMemberType = _channelSyncOrderMemberType
	r.Set("channel_sync_order_member_type", _channelSyncOrderMemberType)
	return nil
}

// GetChannelSyncOrderMemberType ChannelSyncOrderMemberType Getter
func (r AlitripxhotelchannelordermembertypesyncAPIRequest) GetChannelSyncOrderMemberType() *ChannelSyncOrderMemberType {
	return r._channelSyncOrderMemberType
}
