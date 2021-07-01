package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoMerchantContractCancelAPIRequest
蜂鸟商户解约接口 API请求
alibaba.ele.fengniao.merchant.contract.cancel

通过调用此接口，商家及商家下的所有门店解除蜂鸟物流服务 */
type AlibabaEleFengniaoMerchantContractCancelAPIRequest struct {
	model.Params
	// 系统自动生成
	_param *Param
}

// New
