package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstLogisticsSendinfoQueryAPIRequest
供应商-异云-查询主订单包含的物流单 API请求
alibaba.lst.logistics.sendinfo.query

查询主订单包含的物流单 */
type AlibabaLstLogisticsSendinfoQueryAPIRequest struct {
	model.Params
	// 入参
	_query *LstLogisticsInfoQuery
}

// New
