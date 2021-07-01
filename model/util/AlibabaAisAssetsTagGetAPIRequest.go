package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAisAssetsTagGetAPIRequest
基础设施资产标签获取 API请求
alibaba.ais.assets.tag.get

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code获取 */
type AlibabaAisAssetsTagGetAPIRequest struct {
	model.Params
	// 二维码生成唯一标识
	_uNonce string
}

// New
