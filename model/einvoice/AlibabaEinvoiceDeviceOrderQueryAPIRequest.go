package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceDeviceOrderQueryAPIRequest
查询税控设备加盘订购单详情 API请求
alibaba.einvoice.device.order.query

查询税控设备订购单详情 */
type AlibabaEinvoiceDeviceOrderQueryAPIRequest struct {
	model.Params
	// 税控设备订购单ID
	_flowId string
}

// New
