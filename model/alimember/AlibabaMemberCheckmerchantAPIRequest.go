package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMemberCheckmerchantAPIRequest
校验商家身份 API请求
alibaba.member.checkmerchant

校验商家身份 */
type AlibabaMemberCheckmerchantAPIRequest struct {
	model.Params
	// 混淆后的商家id
	_openMerchantId string
}

// New
