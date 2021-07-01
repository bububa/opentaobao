package yunosaccount

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosAccountCallapiAPIRequest
调用YunOS账号开放API API请求
yunos.account.callapi

YunOS账号客户端对外开放的api通过top暴露 */
type YunosAccountCallapiAPIRequest struct {
	model.Params
	// API版本号
	_version string
	// 调用的API名称
	_api string
	// 时间戳，精确到秒；账号服务端会校验该值与服务器当前时间戳的差值，超过一定范围则拒绝请求
	_timeStamp string
	// 业务参数
	_params string
	// 应用签名的MD5值
	_authSign string
}

// New
