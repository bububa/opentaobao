package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMessageListAPIRequest
获取留言列表 API请求
taobao.ailab.aicloud.top.message.list

根据指定参数获取留言列表 */
type TaobaoAilabAicloudTopMessageListAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 截止时间，yyyy-MM-dd HH:mm:ss.SSS
	_beforeTime string
	// 一次性展示多少条信息
	_limit int64
}

// New
