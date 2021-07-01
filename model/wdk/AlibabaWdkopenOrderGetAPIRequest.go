package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkopenOrderGetAPIRequest
五道口商户订单获取 API请求
alibaba.wdkopen.order.get

商户通过五道口订单id获取订单信息 */
type AlibabaWdkopenOrderGetAPIRequest struct {
	model.Params
	// 经营店id
	_storeId string
	// 五道口主订单id
	_bizOrderId int64
	// 外部主订单ID
	_outOrderId string
}

// New
