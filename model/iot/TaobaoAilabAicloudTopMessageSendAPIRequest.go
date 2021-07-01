package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMessageSendAPIRequest
发送留言 API请求
taobao.ailab.aicloud.top.message.send

供准入的外部用户实现发送留言功能，APP端发送，设备端读取 */
type TaobaoAilabAicloudTopMessageSendAPIRequest struct {
	model.Params
	// 扩展信息，用于存放APP类型等
	_ext string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 账户体系隔离
	_schema string
	// 用户上传到OSS上的文件路径，不包含域名
	_audioPath string
}

// New
