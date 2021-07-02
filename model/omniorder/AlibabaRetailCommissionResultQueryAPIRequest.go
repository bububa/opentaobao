package omniorder

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailCommissionResultQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.commission.result.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailCommissionResultQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 请求参数
func (r *AlibabaRetailCommissionResultQueryAPIRequest) SetParam0(_param0 *CommissionResultQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaRetailCommissionResultQueryAPIRequest) GetParam0() *CommissionResultQuery {
	return r._param0
}
