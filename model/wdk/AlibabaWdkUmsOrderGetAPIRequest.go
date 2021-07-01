package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsOrderGetAPIRequest
查询店仓作业单据清单 （库存对账辅助）-回流单 API请求
alibaba.wdk.ums.order.get

查询店仓作业单据清单 （库存对账辅助）-回流单 */
type AlibabaWdkUmsOrderGetAPIRequest struct {
	model.Params
	// 查询单据的dto
	_queryErpbillDto *QueryErpBillDto
}

// New
