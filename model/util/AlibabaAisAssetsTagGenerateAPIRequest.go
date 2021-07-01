package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAisAssetsTagGenerateAPIRequest
基础设施资产标签生成 API请求
alibaba.ais.assets.tag.generate

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code生成 */
type AlibabaAisAssetsTagGenerateAPIRequest struct {
	model.Params
	// 请求资产信息
	_requestParam string
}

// New
