package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopAuthGetAPIRequest
登陆 API请求
taobao.ailab.aicloud.top.auth.get

登陆 */
type TaobaoAilabAicloudTopAuthGetAPIRequest struct {
	model.Params
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 账户体系隔离
	_schema string
	// app类型
	_appType string
}

// New
