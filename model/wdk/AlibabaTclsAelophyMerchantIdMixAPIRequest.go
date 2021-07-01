package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyMerchantIdMixAPIRequest
商家用户id混淆 API请求
alibaba.tcls.aelophy.merchant.id.mix

商家用户id混淆 */
type AlibabaTclsAelophyMerchantIdMixAPIRequest struct {
	model.Params
	// 商家用户id
	_unionUid string
}

// New
