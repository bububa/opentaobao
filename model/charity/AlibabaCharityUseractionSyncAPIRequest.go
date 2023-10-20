package charity

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityUseractionSyncAPIRequest 用户公益行为同步 API请求
// alibaba.charity.useraction.sync
//
// 外部公益活动，用户公益行为同步
type AlibabaCharityUseractionSyncAPIRequest struct {
	model.Params
	// 用户公益行为
	_channelUserActionDto *ChannelUserActionDto
}

// NewAlibabaCharityUseractionSyncRequest 初始化AlibabaCharityUseractionSyncAPIRequest对象
func NewAlibabaCharityUseractionSyncRequest() *AlibabaCharityUseractionSyncAPIRequest {
	return &AlibabaCharityUseractionSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCharityUseractionSyncAPIRequest) Reset() {
	r._channelUserActionDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCharityUseractionSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.useraction.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCharityUseractionSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCharityUseractionSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelUserActionDto is ChannelUserActionDto Setter
// 用户公益行为
func (r *AlibabaCharityUseractionSyncAPIRequest) SetChannelUserActionDto(_channelUserActionDto *ChannelUserActionDto) error {
	r._channelUserActionDto = _channelUserActionDto
	r.Set("channel_user_action_dto", _channelUserActionDto)
	return nil
}

// GetChannelUserActionDto ChannelUserActionDto Getter
func (r AlibabaCharityUseractionSyncAPIRequest) GetChannelUserActionDto() *ChannelUserActionDto {
	return r._channelUserActionDto
}

var poolAlibabaCharityUseractionSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCharityUseractionSyncRequest()
	},
}

// GetAlibabaCharityUseractionSyncRequest 从 sync.Pool 获取 AlibabaCharityUseractionSyncAPIRequest
func GetAlibabaCharityUseractionSyncAPIRequest() *AlibabaCharityUseractionSyncAPIRequest {
	return poolAlibabaCharityUseractionSyncAPIRequest.Get().(*AlibabaCharityUseractionSyncAPIRequest)
}

// ReleaseAlibabaCharityUseractionSyncAPIRequest 将 AlibabaCharityUseractionSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaCharityUseractionSyncAPIRequest(v *AlibabaCharityUseractionSyncAPIRequest) {
	v.Reset()
	poolAlibabaCharityUseractionSyncAPIRequest.Put(v)
}
