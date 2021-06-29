package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店分销渠道会员类型同步 API请求
alitrip.xhotel.channel.order.membertype.sync

酒店分销渠道会员类型同步
*/
type AlitripXhotelChannelOrderMembertypeSyncRequest struct {
    model.Params
    // 入参
    channelSyncOrderMemberType   *ChannelSyncOrderMemberType
}

// 初始化AlitripXhotelChannelOrderMembertypeSyncRequest对象
func NewAlitripXhotelChannelOrderMembertypeSyncRequest() *AlitripXhotelChannelOrderMembertypeSyncRequest{
    return &AlitripXhotelChannelOrderMembertypeSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripXhotelChannelOrderMembertypeSyncRequest) GetApiMethodName() string {
    return "alitrip.xhotel.channel.order.membertype.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlitripXhotelChannelOrderMembertypeSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelSyncOrderMemberType Setter
// 入参
func (r *AlitripXhotelChannelOrderMembertypeSyncRequest) SetChannelSyncOrderMemberType(channelSyncOrderMemberType *ChannelSyncOrderMemberType) error {
    r.channelSyncOrderMemberType = channelSyncOrderMemberType
    r.Set("channel_sync_order_member_type", channelSyncOrderMemberType)
    return nil
}

// ChannelSyncOrderMemberType Getter
func (r AlitripXhotelChannelOrderMembertypeSyncRequest) GetChannelSyncOrderMemberType() *ChannelSyncOrderMemberType {
    return r.channelSyncOrderMemberType
}
