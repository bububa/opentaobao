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

// NewAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest 初始化AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest对象
func NewAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest() *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest {
	return &AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.advance.bill.flow.direction"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest) GetCode() string {
	return r._code
}
