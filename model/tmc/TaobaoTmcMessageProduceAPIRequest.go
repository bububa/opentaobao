package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcMessageProduceAPIRequest
发布单条消息 API请求
taobao.tmc.message.produce

发布单条消息 */
type TaobaoTmcMessageProduceAPIRequest struct {
	model.Params
	// 消息内容的JSON表述，必须按照topic的定义来填充
	_content string
	// 消息类型
	_topic string
	// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。
	_mediaContent *model.File
	// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
	_mediaContent2 *model.File
	// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
	_mediaContent3 *model.File
	// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
	_mediaContent4 *model.File
	// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
	_mediaContent5 *model.File
	// 目标分组，一般为default
	_targetGroup string
}

// NewTaobaoTmcMessageProduceRequest 初始化TaobaoTmcMessageProduceAPIRequest对象
func NewTaobaoTmcMessageProduceRequest() *TaobaoTmcMessageProduceAPIRequest {
	return &TaobaoTmcMessageProduceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcMessageProduceAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.message.produce"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcMessageProduceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Content Setter
// 消息内容的JSON表述，必须按照topic的定义来填充
func (r *TaobaoTmcMessageProduceAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// Get Content Getter
func (r TaobaoTmcMessageProduceAPIRequest) GetContent() string {
	return r._content
}

// Set is Topic Setter
// 消息类型
func (r *TaobaoTmcMessageProduceAPIRequest) SetTopic(_topic string) error {
	r._topic = _topic
	r.Set("topic", _topic)
	return nil
}

// Get Topic Getter
func (r TaobaoTmcMessageProduceAPIRequest) GetTopic() string {
	return r._topic
}

// Set is MediaContent Setter
// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。
func (r *TaobaoTmcMessageProduceAPIRequest) SetMediaContent(_mediaContent *model.File) error {
	r._mediaContent = _mediaContent
	r.Set("media_content", _mediaContent)
	return nil
}

// Get MediaContent Getter
func (r TaobaoTmcMessageProduceAPIRequest) GetMediaContent() *model.File {
	return r._mediaContent
}

// Set is MediaContent2 Setter
// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
func (r *TaobaoTmcMessageProduceAPIRequest) SetMediaContent2(_mediaContent2 *model.File) error {
	r._mediaContent2 = _mediaContent2
	r.Set("media_content2", _mediaContent2)
	return nil
}

// Get MediaContent2 Getter
func (r TaobaoTmcMessageProduceAPIRequest) GetMediaContent2() *model.File {
	return r._mediaContent2
}

// Set is MediaContent3 Setter
// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
func (r *TaobaoTmcMessageProduceAPIRequest) SetMediaContent3(_mediaContent3 *model.File) error {
	r._mediaContent3 = _mediaContent3
	r.Set("media_content3", _mediaContent3)
	return nil
}

// Get MediaContent3 Getter
func (r TaobaoTmcMessageProduceAPIRequest) GetMediaContent3() *model.File {
	return r._mediaContent3
}

// Set is MediaContent4 Setter
// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
func (r *TaobaoTmcMessageProduceAPIRequest) SetMediaContent4(_mediaContent4 *model.File) error {
	r._mediaContent4 = _mediaContent4
	r.Set("media_content4", _mediaContent4)
	return nil
}

// Get MediaContent4 Getter
func (r TaobaoTmcMessageProduceAPIRequest) GetMediaContent4() *model.File {
	return r._mediaContent4
}

// Set is MediaContent5 Setter
// 回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。
func (r *TaobaoTmcMessageProduceAPIRequest) SetMediaContent5(_mediaContent5 *model.File) error {
	r._mediaContent5 = _mediaContent5
	r.Set("media_content5", _mediaContent5)
	return nil
}

// Get MediaContent5 Getter
func (r TaobaoTmcMessageProduceAPIRequest) GetMediaContent5() *model.File {
	return r._mediaContent5
}

// Set is TargetGroup Setter
// 目标分组，一般为default
func (r *TaobaoTmcMessageProduceAPIRequest) SetTargetGroup(_targetGroup string) error {
	r._targetGroup = _targetGroup
	r.Set("target_group", _targetGroup)
	return nil
}

// Get TargetGroup Getter
func (r TaobaoTmcMessageProduceAPIRequest) GetTargetGroup() string {
	return r._targetGroup
}
