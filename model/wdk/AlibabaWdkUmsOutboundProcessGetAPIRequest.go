package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsOutboundProcessGetAPIRequest
出库业务UMS异步处理结果返回 API请求
alibaba.wdk.ums.outbound.process.get

出库业务UMS异步处理结果返回 */
type AlibabaWdkUmsOutboundProcessGetAPIRequest struct {
	model.Params
	// 店仓code，指的是作业对象，对应一个物理店或仓编码
	_warehouseCode string
}

// New
