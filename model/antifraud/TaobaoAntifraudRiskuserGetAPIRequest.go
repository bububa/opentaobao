package antifraud

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
反欺诈用户风险查询 API请求
taobao.antifraud.riskuser.get

根据用户基础信息，核实平台上的用户是否存在欺诈风险
*/
type TaobaoAntifraudRiskuserGetAPIRequest struct {
    model.Params
    // 风险用户查询条件
    _paramAccountQuery   *ParamAccountQuery
}

// 初始化TaobaoAntifraudRiskuserGetAPIRequest对象
func NewTaobaoAntifraudRiskuserGetRequest() *TaobaoAntifraudRiskuserGetAPIRequest{
    return &TaobaoAntifraudRiskuserGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAntifraudRiskuserGetAPIRequest) GetApiMethodName() string {
    return "taobao.antifraud.riskuser.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAntifraudRiskuserGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamAccountQuery Setter
// 风险用户查询条件
func (r *TaobaoAntifraudRiskuserGetAPIRequest) SetParamAccountQuery(_paramAccountQuery *ParamAccountQuery) error {
    r._paramAccountQuery = _paramAccountQuery
    r.Set("param_account_query", _paramAccountQuery)
    return nil
}

// ParamAccountQuery Getter
func (r TaobaoAntifraudRiskuserGetAPIRequest) GetParamAccountQuery() *ParamAccountQuery {
    return r._paramAccountQuery
}
