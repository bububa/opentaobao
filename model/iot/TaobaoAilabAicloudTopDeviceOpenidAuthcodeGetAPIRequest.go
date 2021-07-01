package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest
获取openid设备通用授权码 API请求
taobao.ailab.aicloud.top.device.openid.authcode.get

获取openid设备通用授权码 */
type TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest struct {
	model.Params
	// 淘宝openid
	_openId string
	// 账户体系隔离，即硬件接入平台中取得的schema key。
	_schema string
	// (废弃) 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
}

// New
