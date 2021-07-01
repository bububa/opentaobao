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
type AlibabaCharityUseractionSyncAPIRequest struct {
    model.Params
    // 用户公益行为
    _channelUserActionDto   *ChannelUserActionDto
}

// 初始化AlibabaCharityUseractionSyncAPIRequest对象
func NewAlibabaCharityUseractionSyncRequest() *AlibabaCharityUseractionSyncAPIRequest{
    return &AlibabaCharityUseractionSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCharityUseractionSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.charity.useraction.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCharityUseractionSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelUserActionDto Setter
// 用户公益行为
func (r *AlibabaCharityUseractionSyncAPIRequest) SetChannelUserActionDto(_channelUserActionDto *ChannelUserActionDto) error {
    r._channelUserActionDto = _channelUserActionDto
    r.Set("channel_user_action_dto", _channelUserActionDto)
    return nil
}

// ChannelUserActionDto Getter
func (r AlibabaCharityUseractionSyncAPIRequest) GetChannelUserActionDto() *ChannelUserActionDto {
    return r._channelUserActionDto
}
