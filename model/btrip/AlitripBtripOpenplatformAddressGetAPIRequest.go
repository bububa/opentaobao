package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripOpenplatformAddressGetAPIRequest
【商旅】开放平台对外页面跳转 API请求
alitrip.btrip.openplatform.address.get

获取类目预定页跳转地址 */
type AlitripBtripOpenplatformAddressGetAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiJumpInfoRq
}

// New
