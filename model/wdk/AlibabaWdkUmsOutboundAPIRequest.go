package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsOutboundAPIRequest
出库-ERP下发单(新接口，包含调拨出库单和退货出库单等) API请求
alibaba.wdk.ums.outbound

出库-ERP下发单(新接口，包含调拨出库单和退货出库单等) */
type AlibabaWdkUmsOutboundAPIRequest struct {
	model.Params
	// 出库-ERP下发单请求dto
	_erpOutputOrderDto *ErpOutputOrderDto
}

// New
