package antifraud

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoantifraudriskusergetAPIRequest 反欺诈用户风险查询 API请求
// taobao.antifraud.riskuser.get
//
// 根据用户基础信息，核实平台上的用户是否存在欺诈风险
type TaobaoantifraudriskusergetAPIRequest struct {
	model.Params
	// 风险用户查询条件
	_paramAccountQuery *ParamAccountQuery
}

// NewTaobaoantifraudriskusergetRequest 初始化TaobaoantifraudriskusergetAPIRequest对象
func NewTaobaoantifraudriskusergetRequest() *TaobaoantifraudriskusergetAPIRequest {
	return &TaobaoantifraudriskusergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoantifraudriskusergetAPIRequest) GetApiMethodName() string {
	return "taobao.antifraud.riskuser.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoantifraudriskusergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoantifraudriskusergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAccountQuery is ParamAccountQuery Setter
// 风险用户查询条件
func (r *TaobaoantifraudriskusergetAPIRequest) SetParamAccountQuery(_paramAccountQuery *ParamAccountQuery) error {
	r._paramAccountQuery = _paramAccountQuery
	r.Set("param_account_query", _paramAccountQuery)
	return nil
}

// GetParamAccountQuery ParamAccountQuery Getter
func (r TaobaoantifraudriskusergetAPIRequest) GetParamAccountQuery() *ParamAccountQuery {
	return r._paramAccountQuery
}
