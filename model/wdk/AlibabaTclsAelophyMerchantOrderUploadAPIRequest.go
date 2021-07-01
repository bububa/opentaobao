package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyMerchantOrderUploadAPIRequest
商家订单数据上传 API请求
alibaba.tcls.aelophy.merchant.order.upload

商家订单数据上传 */
type AlibabaTclsAelophyMerchantOrderUploadAPIRequest struct {
	model.Params
	// 商家订单信息
	_orderInfo *MerchantOrderInfo
}

// New
