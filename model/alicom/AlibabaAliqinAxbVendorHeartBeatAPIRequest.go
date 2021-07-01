package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinAxbVendorHeartBeatAPIRequest
供应商心跳上报接口 API请求
alibaba.aliqin.axb.vendor.heart.beat

供应商上报自己的心跳信息 */
type AlibabaAliqinAxbVendorHeartBeatAPIRequest struct {
	model.Params
	// 可选的预留字段
	_status string
	// 供应商合作KEY
	_vendorKey string
}

// New
