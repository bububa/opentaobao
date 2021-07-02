package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTxdCrmStatementBackflowAPIRequest 淘鲜达商家会员账单回流 API请求
// alibaba.wdk.txd.crm.statement.backflow
//
// 淘鲜达商家会员账单回流
type AlibabaWdkTxdCrmStatementBackflowAPIRequest struct {
	model.Params
	// 参数
	_paramStatementBO *StatementBo
}

// NewAlibabaWdkTxdCrmStatementBackflowRequest 初始化AlibabaWdkTxdCrmStatementBackflowAPIRequest对象
func NewAlibabaWdkTxdCrmStatementBackflowRequest() *AlibabaWdkTxdCrmStatementBackflowAPIRequest {
	return &AlibabaWdkTxdCrmStatementBackflowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTxdCrmStatementBackflowAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.txd.crm.statement.backflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTxdCrmStatementBackflowAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamStatementBO Setter
// 参数
func (r *AlibabaWdkTxdCrmStatementBackflowAPIRequest) SetParamStatementBO(_paramStatementBO *StatementBo) error {
	r._paramStatementBO = _paramStatementBO
	r.Set("param_statement_b_o", _paramStatementBO)
	return nil
}

// Get ParamStatementBO Getter
func (r AlibabaWdkTxdCrmStatementBackflowAPIRequest) GetParamStatementBO() *StatementBo {
	return r._paramStatementBO
}
