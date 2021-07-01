package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkorderSharestockInsuranceCallbackAPIRequest
共享库存订单投保后回传保单号 API请求
alibaba.wdkorder.sharestock.insurance.callback

共享库存订单投保消息获取 */
type AlibabaWdkorderSharestockInsuranceCallbackAPIRequest struct {
	model.Params
	// 淘宝交易子订单ID
	_tbSubOrderId int64
	// 投保单ID
	_insuranceId string
}

// New
