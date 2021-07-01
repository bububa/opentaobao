package antifraud

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAntifraudRiskuserGetAPIRequest
反欺诈用户风险查询 API请求
taobao.antifraud.riskuser.get

根据用户基础信息，核实平台上的用户是否存在欺诈风险 */
type TaobaoAntifraudRiskuserGetAPIRequest struct {
	model.Params
	// 风险用户查询条件
	_paramAccountQuery *ParamAccountQuery
}

// New
