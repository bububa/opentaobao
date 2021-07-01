package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjOcWritesaleslipAPIRequest
开票占库 API请求
alibaba.mj.oc.writesaleslip

开票占库 */
type AlibabaMjOcWritesaleslipAPIRequest struct {
	model.Params
	// 开票占库入参
	_posSaleOrder *PosSaleOrderDto
}

// New
