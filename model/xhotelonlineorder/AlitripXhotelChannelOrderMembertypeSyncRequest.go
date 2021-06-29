package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店分销渠道会员类型同步 APIRequest
alitrip.xhotel.channel.order.membertype.sync

酒店分销渠道会员类型同步
*/
type AlitripXhotelChannelOrderMembertypeSyncRequest struct {
    model.Params

    // 入参
    channelSyncOrderMemberType   *ChannelSyncOrderMemberType 

}

func NewAlitripXhotelChannelOrderMembertypeSyncRequest() *AlitripXhotelChannelOrderMembertypeSyncRequest{
    return &AlitripXhotelChannelOrderMembertypeSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripXhotelChannelOrderMembertypeSyncRequest) GetApiMethodName() string {
    return "alitrip.xhotel.channel.order.membertype.sync"
}

func (r AlitripXhotelChannelOrderMembertypeSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripXhotelChannelOrderMembertypeSyncRequest) SetChannelSyncOrderMemberType(channelSyncOrderMemberType *ChannelSyncOrderMemberType) error {
    r.channelSyncOrderMemberType = channelSyncOrderMemberType
    r.Set("channel_sync_order_member_type", channelSyncOrderMemberType)
    return nil
}

func (r AlitripXhotelChannelOrderMembertypeSyncRequest) GetChannelSyncOrderMemberType() *ChannelSyncOrderMemberType {
    return r.channelSyncOrderMemberType
}

