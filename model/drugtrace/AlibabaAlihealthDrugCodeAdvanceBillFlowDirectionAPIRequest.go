package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest 单据流向查询 API请求
// alibaba.alihealth.drug.code.advance.bill.flow.direction
//
// 单据流向查询
type AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest struct {
	model.Params
	// 追溯码
	_code string
}

// NewAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest 初始化AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest对象
func NewAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest() *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest {
	return &AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest) Reset() {
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.advance.bill.flow.direction"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest) GetCode() string {
	return r._code
}

var poolAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest()
	},
}

// GetAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionRequest 从 sync.Pool 获取 AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest
func GetAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest() *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest {
	return poolAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest.Get().(*AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest)
}

// ReleaseAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest 将 AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest(v *AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest.Put(v)
}
