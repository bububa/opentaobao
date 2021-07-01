package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMessageSendtextAPIRequest
故事机发送文本留言 API请求
taobao.ailab.aicloud.top.message.sendtext

故事机文本留言 */
type TaobaoAilabAicloudTopMessageSendtextAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户 id
	_userId string
	// 用户设备唯一识别码，长度限制32以内， 建议使用系统接口获取deviceid, 然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 留言输入的文本
	_text string
}

// New
