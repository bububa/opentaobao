package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstLogisticsTraceQueryAPIRequest
供应商-异云-查询运单物流追踪信息 API请求
alibaba.lst.logistics.trace.query

查询LP单物流追踪信息 */
type AlibabaLstLogisticsTraceQueryAPIRequest struct {
	model.Params
	// 入参
	_query *LstLogisticsTraceQuery
}

// New
