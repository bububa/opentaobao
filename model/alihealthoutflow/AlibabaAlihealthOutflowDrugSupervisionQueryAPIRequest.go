package alihealthoutflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest 监管平台药品查询 API请求
// alibaba.alihealth.outflow.drug.supervision.query
//
// 获取监管平台药品数据
type AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest struct {
	model.Params
	// 请求
	_request1 *OuterDrugVo
}

// NewAlibabaAlihealthOutflowDrugSupervisionQueryRequest 初始化AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest对象
func NewAlibabaAlihealthOutflowDrugSupervisionQueryRequest() *AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest {
	return &AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest) Reset() {
	r._request1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.outflow.drug.supervision.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest1 is Request1 Setter
// 请求
func (r *AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest) SetRequest1(_request1 *OuterDrugVo) error {
	r._request1 = _request1
	r.Set("request1", _request1)
	return nil
}

// GetRequest1 Request1 Getter
func (r AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest) GetRequest1() *OuterDrugVo {
	return r._request1
}

var poolAlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthOutflowDrugSupervisionQueryRequest()
	},
}

// GetAlibabaAlihealthOutflowDrugSupervisionQueryRequest 从 sync.Pool 获取 AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest
func GetAlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest() *AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest {
	return poolAlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest.Get().(*AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest)
}

// ReleaseAlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest 将 AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest(v *AlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthOutflowDrugSupervisionQueryAPIRequest.Put(v)
}
