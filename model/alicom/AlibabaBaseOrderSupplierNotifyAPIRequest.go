package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBaseOrderSupplierNotifyAPIRequest
阿里通信运营商信息回传 API请求
alibaba.base.order.supplier.notify

接收阿里通信流量运营商信息回传 */
type AlibabaBaseOrderSupplierNotifyAPIRequest struct {
	model.Params
	// 入参对象
	_paramFlowSuppllierNotifyModel *FlowSuppllierNotifyModel
}

// New
