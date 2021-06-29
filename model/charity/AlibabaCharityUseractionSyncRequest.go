package charity

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户公益行为同步 API请求
alibaba.charity.useraction.sync

外部公益活动，用户公益行为同步
*/
type AlibabaCharityUseractionSyncRequest struct {
    model.Params
    // 用户公益行为
    channelUserActionDto   *ChannelUserActionDto
}

// 初始化AlibabaCharityUseractionSyncRequest对象
func NewAlibabaCharityUseractionSyncRequest() *AlibabaCharityUseractionSyncRequest{
    return &AlibabaCharityUseractionSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCharityUseractionSyncRequest) GetApiMethodName() string {
    return "alibaba.charity.useraction.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCharityUseractionSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelUserActionDto Setter
// 用户公益行为
func (r *AlibabaCharityUseractionSyncRequest) SetChannelUserActionDto(channelUserActionDto *ChannelUserActionDto) error {
    r.channelUserActionDto = channelUserActionDto
    r.Set("channel_user_action_dto", channelUserActionDto)
    return nil
}

// ChannelUserActionDto Getter
func (r AlibabaCharityUseractionSyncRequest) GetChannelUserActionDto() *ChannelUserActionDto {
    return r.channelUserActionDto
}
