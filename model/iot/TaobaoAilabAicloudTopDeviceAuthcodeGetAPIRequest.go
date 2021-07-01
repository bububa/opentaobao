package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest
获取设备授权码 API请求
taobao.ailab.aicloud.top.device.authcode.get

获取设备授权码 */
type TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest struct {
	model.Params
	// 账户体系隔离，即硬件接入平台中取得的schema key。
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id，由开发者或厂商自行定义，每一schema key下唯一即可
	_userId string
	// (废弃) 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
}

// New
