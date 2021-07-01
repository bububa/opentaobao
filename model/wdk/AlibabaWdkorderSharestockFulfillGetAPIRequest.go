package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkorderSharestockFulfillGetAPIRequest
商户订单履约数据获取 API请求
alibaba.wdkorder.sharestock.fulfill.get

商户订单履约数据获取 */
type AlibabaWdkorderSharestockFulfillGetAPIRequest struct {
	model.Params
	// 履约单ID
	_fulfillOrderId string
}

// New
