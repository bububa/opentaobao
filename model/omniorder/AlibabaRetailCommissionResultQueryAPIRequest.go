package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailCommissionResultQueryAPIRequest 分佣结果查询 API请求
// alibaba.retail.commission.result.query
//
// 查询导购分佣记录
type AlibabaRetailCommissionResultQueryAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *CommissionResultQuery
}

// NewAlibabaRetailCommissionResultQueryRequest 初始化AlibabaRetailCommissionResultQueryAPIRequest对象
func NewAlibabaRetailCommissionResultQueryRequest() *AlibabaRetailCommissionResultQueryAPIRequest {
	return &AlibabaRetailCommissionResultQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailCommissionResultQueryAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailCommissionResultQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.commission.result.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailCommissionResultQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailCommissionResultQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求参数
func (r *AlibabaRetailCommissionResultQueryAPIRequest) SetParam0(_param0 *CommissionResultQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaRetailCommissionResultQueryAPIRequest) GetParam0() *CommissionResultQuery {
	return r._param0
}

var poolAlibabaRetailCommissionResultQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailCommissionResultQueryRequest()
	},
}

// GetAlibabaRetailCommissionResultQueryRequest 从 sync.Pool 获取 AlibabaRetailCommissionResultQueryAPIRequest
func GetAlibabaRetailCommissionResultQueryAPIRequest() *AlibabaRetailCommissionResultQueryAPIRequest {
	return poolAlibabaRetailCommissionResultQueryAPIRequest.Get().(*AlibabaRetailCommissionResultQueryAPIRequest)
}

// ReleaseAlibabaRetailCommissionResultQueryAPIRequest 将 AlibabaRetailCommissionResultQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailCommissionResultQueryAPIRequest(v *AlibabaRetailCommissionResultQueryAPIRequest) {
	v.Reset()
	poolAlibabaRetailCommissionResultQueryAPIRequest.Put(v)
}
