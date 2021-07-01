package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMessageSendaudioAPIRequest
发送语音留言 API请求
taobao.ailab.aicloud.top.message.sendaudio

将语音的二进制byte[]通过TOP接口发送保存 */
type TaobaoAilabAicloudTopMessageSendaudioAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户 id
	_userId string
	// 用户设备唯一识别码，长度限制32以内， 建议使用系统接口获取deviceid, 然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 语音的二进制
	_message *model.File
}

// New
