package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkorderSharestockInsuranceGetorderAPIRequest
共享库存订单投保消息获取 API请求
alibaba.wdkorder.sharestock.insurance.getorder

共享库存订单投保消息获取 */
type AlibabaWdkorderSharestockInsuranceGetorderAPIRequest struct {
	model.Params
	// 淘宝子订单ID
	_tbSubOrderId int64
}

// New
