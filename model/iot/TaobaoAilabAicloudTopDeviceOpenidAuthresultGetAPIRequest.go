package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest
获取openId设备授权码验证结果 API请求
taobao.ailab.aicloud.top.device.openid.authresult.get

获取openId设备授权码验证结果 */
type TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest struct {
	model.Params
	// 淘宝openid
	_openId string
	// 账户体系隔离
	_schema string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// authcode list
	_authCodes []string
}

// New
