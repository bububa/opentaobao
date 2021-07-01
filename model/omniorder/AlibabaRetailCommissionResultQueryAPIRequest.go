package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailCommissionResultQueryAPIRequest
分佣结果查询 API请求
alibaba.retail.commission.result.query

查询导购分佣记录 */
type AlibabaRetailCommissionResultQueryAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *CommissionResultQuery
}

// New
