package antifraud

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAntifraudRiskuserGetAPIRequest 反欺诈用户风险查询 API请求
// taobao.antifraud.riskuser.get
//
// 根据用户基础信息，核实平台上的用户是否存在欺诈风险
type TaobaoAntifraudRiskuserGetAPIRequest struct {
	model.Params
	// 风险用户查询条件
	_paramAccountQuery *ParamAccountQuery
}

// NewTaobaoAntifraudRiskuserGetRequest 初始化TaobaoAntifraudRiskuserGetAPIRequest对象
func NewTaobaoAntifraudRiskuserGetRequest() *TaobaoAntifraudRiskuserGetAPIRequest {
	return &TaobaoAntifraudRiskuserGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAntifraudRiskuserGetAPIRequest) Reset() {
	r._paramAccountQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAntifraudRiskuserGetAPIRequest) GetApiMethodName() string {
	return "taobao.antifraud.riskuser.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAntifraudRiskuserGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAntifraudRiskuserGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAccountQuery is ParamAccountQuery Setter
// 风险用户查询条件
func (r *TaobaoAntifraudRiskuserGetAPIRequest) SetParamAccountQuery(_paramAccountQuery *ParamAccountQuery) error {
	r._paramAccountQuery = _paramAccountQuery
	r.Set("param_account_query", _paramAccountQuery)
	return nil
}

// GetParamAccountQuery ParamAccountQuery Getter
func (r TaobaoAntifraudRiskuserGetAPIRequest) GetParamAccountQuery() *ParamAccountQuery {
	return r._paramAccountQuery
}

var poolTaobaoAntifraudRiskuserGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAntifraudRiskuserGetRequest()
	},
}

// GetTaobaoAntifraudRiskuserGetRequest 从 sync.Pool 获取 TaobaoAntifraudRiskuserGetAPIRequest
func GetTaobaoAntifraudRiskuserGetAPIRequest() *TaobaoAntifraudRiskuserGetAPIRequest {
	return poolTaobaoAntifraudRiskuserGetAPIRequest.Get().(*TaobaoAntifraudRiskuserGetAPIRequest)
}

// ReleaseTaobaoAntifraudRiskuserGetAPIRequest 将 TaobaoAntifraudRiskuserGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAntifraudRiskuserGetAPIRequest(v *TaobaoAntifraudRiskuserGetAPIRequest) {
	v.Reset()
	poolTaobaoAntifraudRiskuserGetAPIRequest.Put(v)
}
