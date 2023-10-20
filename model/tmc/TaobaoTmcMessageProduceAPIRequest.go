package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmcmessageproduceAPIRequest 发布单条消息 API请求
// taobao.tmc.message.produce
//
// 发布单条消息
type TaobaotmcmessageproduceAPIRequest struct {
	model.Params
	// 消息内容的JSON表述，必须按照topic的定义来填充
	_content string
	// 目标分组，一般为default
	_targetGroup string
	// 消息类型
	_topic string
	// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。
	_mediaContent *model.File
	// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
	_mediaContent2 *model.File
	// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
	_mediaContent3 *model.File
	// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
	_mediaContent5 *model.File
	// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
	_mediaContent4 *model.File
	// 延时参数 时间戳 预期发送时间
	_delayMillis int64
	// 提前过期 相对时间差 毫秒
	_expiresMillis int64
}

// NewTaobaotmcmessageproduceRequest 初始化TaobaotmcmessageproduceAPIRequest对象
func NewTaobaotmcmessageproduceRequest() *TaobaotmcmessageproduceAPIRequest {
	return &TaobaotmcmessageproduceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotmcmessageproduceAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.message.produce"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotmcmessageproduceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotmcmessageproduceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 消息内容的JSON表述，必须按照topic的定义来填充
func (r *TaobaotmcmessageproduceAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaotmcmessageproduceAPIRequest) GetContent() string {
	return r._content
}

// SetTargetGroup is TargetGroup Setter
// 目标分组，一般为default
func (r *TaobaotmcmessageproduceAPIRequest) SetTargetGroup(_targetGroup string) error {
	r._targetGroup = _targetGroup
	r.Set("target_group", _targetGroup)
	return nil
}

// GetTargetGroup TargetGroup Getter
func (r TaobaotmcmessageproduceAPIRequest) GetTargetGroup() string {
	return r._targetGroup
}

// SetTopic is Topic Setter
// 消息类型
func (r *TaobaotmcmessageproduceAPIRequest) SetTopic(_topic string) error {
	r._topic = _topic
	r.Set("topic", _topic)
	return nil
}

// GetTopic Topic Getter
func (r TaobaotmcmessageproduceAPIRequest) GetTopic() string {
	return r._topic
}

// SetMediaContent is MediaContent Setter
// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。
func (r *TaobaotmcmessageproduceAPIRequest) SetMediaContent(_mediaContent *model.File) error {
	r._mediaContent = _mediaContent
	r.Set("media_content", _mediaContent)
	return nil
}

// GetMediaContent MediaContent Getter
func (r TaobaotmcmessageproduceAPIRequest) GetMediaContent() *model.File {
	return r._mediaContent
}

// SetMediaContent2 is MediaContent2 Setter
// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
func (r *TaobaotmcmessageproduceAPIRequest) SetMediaContent2(_mediaContent2 *model.File) error {
	r._mediaContent2 = _mediaContent2
	r.Set("media_content2", _mediaContent2)
	return nil
}

// GetMediaContent2 MediaContent2 Getter
func (r TaobaotmcmessageproduceAPIRequest) GetMediaContent2() *model.File {
	return r._mediaContent2
}

// SetMediaContent3 is MediaContent3 Setter
// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
func (r *TaobaotmcmessageproduceAPIRequest) SetMediaContent3(_mediaContent3 *model.File) error {
	r._mediaContent3 = _mediaContent3
	r.Set("media_content3", _mediaContent3)
	return nil
}

// GetMediaContent3 MediaContent3 Getter
func (r TaobaotmcmessageproduceAPIRequest) GetMediaContent3() *model.File {
	return r._mediaContent3
}

// SetMediaContent5 is MediaContent5 Setter
// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
func (r *TaobaotmcmessageproduceAPIRequest) SetMediaContent5(_mediaContent5 *model.File) error {
	r._mediaContent5 = _mediaContent5
	r.Set("media_content5", _mediaContent5)
	return nil
}

// GetMediaContent5 MediaContent5 Getter
func (r TaobaotmcmessageproduceAPIRequest) GetMediaContent5() *model.File {
	return r._mediaContent5
}

// SetMediaContent4 is MediaContent4 Setter
// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
func (r *TaobaotmcmessageproduceAPIRequest) SetMediaContent4(_mediaContent4 *model.File) error {
	r._mediaContent4 = _mediaContent4
	r.Set("media_content4", _mediaContent4)
	return nil
}

// GetMediaContent4 MediaContent4 Getter
func (r TaobaotmcmessageproduceAPIRequest) GetMediaContent4() *model.File {
	return r._mediaContent4
}

// SetDelayMillis is DelayMillis Setter
// 延时参数 时间戳 预期发送时间
func (r *TaobaotmcmessageproduceAPIRequest) SetDelayMillis(_delayMillis int64) error {
	r._delayMillis = _delayMillis
	r.Set("delay_millis", _delayMillis)
	return nil
}

// GetDelayMillis DelayMillis Getter
func (r TaobaotmcmessageproduceAPIRequest) GetDelayMillis() int64 {
	return r._delayMillis
}

// SetExpiresMillis is ExpiresMillis Setter
// 提前过期 相对时间差 毫秒
func (r *TaobaotmcmessageproduceAPIRequest) SetExpiresMillis(_expiresMillis int64) error {
	r._expiresMillis = _expiresMillis
	r.Set("expires_millis", _expiresMillis)
	return nil
}

// GetExpiresMillis ExpiresMillis Getter
func (r TaobaotmcmessageproduceAPIRequest) GetExpiresMillis() int64 {
	return r._expiresMillis
}
