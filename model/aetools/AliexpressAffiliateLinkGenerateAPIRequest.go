package aetools

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressAffiliateLinkGenerateAPIRequest
联盟推广链接生成 API请求
aliexpress.affiliate.link.generate

AE联盟推广链接生成接口 */
type AliexpressAffiliateLinkGenerateAPIRequest struct {
	model.Params
	// API请求签名
	_appSignature string
	// 转换的链接类型：0代表普通Link，1代表Search Link，2代表 hot link
	_promotionLinkType int64
	// 原始链接或者值
	_sourceValues string
	// 推广者原始trackingID
	_trackingId string
}

// New
