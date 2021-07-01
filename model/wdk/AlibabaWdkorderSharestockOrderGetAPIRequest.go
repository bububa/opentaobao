package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkorderSharestockOrderGetAPIRequest
猫超商户订单拉取 API请求
alibaba.wdkorder.sharestock.order.get

商户拉取猫超订单数据 */
type AlibabaWdkorderSharestockOrderGetAPIRequest struct {
	model.Params
	// 淘宝主订单ID
	_tbOrderId int64
}

// New
