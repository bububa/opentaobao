package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkTxdCrmStatementBackflowAPIRequest
淘鲜达商家会员账单回流 API请求
alibaba.wdk.txd.crm.statement.backflow

淘鲜达商家会员账单回流 */
type AlibabaWdkTxdCrmStatementBackflowAPIRequest struct {
	model.Params
	// 参数
	_paramStatementBO *StatementBo
}

// New
