package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailcommissionresultqueryAPIRequest 分佣结果查询 API请求
// alibaba.retail.commission.result.query
//
// 查询导购分佣记录
type AlibabaretailcommissionresultqueryAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *CommissionResultQuery
}

// NewAlibabaretailcommissionresultqueryRequest 初始化AlibabaretailcommissionresultqueryAPIRequest对象
func NewAlibabaretailcommissionresultqueryRequest() *AlibabaretailcommissionresultqueryAPIRequest {
	return &AlibabaretailcommissionresultqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailcommissionresultqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.commission.result.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailcommissionresultqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailcommissionresultqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求参数
func (r *AlibabaretailcommissionresultqueryAPIRequest) SetParam0(_param0 *CommissionResultQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaretailcommissionresultqueryAPIRequest) GetParam0() *CommissionResultQuery {
	return r._param0
}
