package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsInboundAPIRequest
入库-ERP下发单 API请求
alibaba.wdk.ums.inbound

入库-ERP下发单 */
type AlibabaWdkUmsInboundAPIRequest struct {
	model.Params
	// 1
	_erpArrivalnoticeDto *ErpArrivalNoticeDto
}

// New
