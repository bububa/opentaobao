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

// New
