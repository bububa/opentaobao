package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabUserTokenGetAPIRequest
三方账号获取 token API请求
alibaba.ailab.user.token.get

inside 设备的三方 app，通过 extId、schema 生成 token */
type AlibabaAilabUserTokenGetAPIRequest struct {
	model.Params
	// 三方用户的唯一ID
	_merchantUserId string
	// 开放平台申请的schema
	_schemaKey string
	// 用户点击同意授权，则会有授权结果：success/fail，此结果通过 callBackUrl 回调给三方 如果授权账号重复授权给已授权的淘宝账号，幂等返回成功 url 的调用是 表单 post 的方式， request body success example: merchantUserId=xxx&result=success request body fail example: merchantUserId=xxx&result=fail
	_callBackUrl string
}

// New
