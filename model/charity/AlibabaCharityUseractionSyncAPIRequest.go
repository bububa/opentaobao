package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharityuseractionsyncAPIRequest 用户公益行为同步 API请求
// alibaba.charity.useraction.sync
//
// 外部公益活动，用户公益行为同步
type AlibabacharityuseractionsyncAPIRequest struct {
	model.Params
	// 用户公益行为
	_channelUserActionDto *ChannelUserActionDto
}

// NewAlibabacharityuseractionsyncRequest 初始化AlibabacharityuseractionsyncAPIRequest对象
func NewAlibabacharityuseractionsyncRequest() *AlibabacharityuseractionsyncAPIRequest {
	return &AlibabacharityuseractionsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacharityuseractionsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.useraction.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacharityuseractionsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacharityuseractionsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannelUserActionDto is ChannelUserActionDto Setter
// 用户公益行为
func (r *AlibabacharityuseractionsyncAPIRequest) SetChannelUserActionDto(_channelUserActionDto *ChannelUserActionDto) error {
	r._channelUserActionDto = _channelUserActionDto
	r.Set("channel_user_action_dto", _channelUserActionDto)
	return nil
}

// GetChannelUserActionDto ChannelUserActionDto Getter
func (r AlibabacharityuseractionsyncAPIRequest) GetChannelUserActionDto() *ChannelUserActionDto {
	return r._channelUserActionDto
}
