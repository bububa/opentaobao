package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtPayMerchantStallSigningModifyAPIRequest
三级商户进件修改 API请求
tmall.nrt.pay.merchant.stall.signing.modify

三级商户进件修改 */
type TmallNrtPayMerchantStallSigningModifyAPIRequest struct {
	model.Params
	// 请求参数
	_req *StallSigningReqDto
}

// New
