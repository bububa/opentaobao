package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest
单据流向查询 API请求
alibaba.alihealth.drug.code.advance.bill.flow.direction

单据流向查询 */
type AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest struct {
	model.Params
	// 追溯码
	_code string
}

// New
