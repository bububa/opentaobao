package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAisAssetsTagAbortAPIRequest
基础设施资产标签废弃 API请求
alibaba.ais.assets.tag.abort

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code未使用的废弃 */
type AlibabaAisAssetsTagAbortAPIRequest struct {
	model.Params
	// 请求资产信息
	_requestParam string
}

// New
