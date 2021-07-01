package antifraud

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/antifraud"
)

/* TaobaoAntifraudRiskuserGet
反欺诈用户风险查询
taobao.antifraud.riskuser.get

根据用户基础信息，核实平台上的用户是否存在欺诈风险 */
func TaobaoAntifraudRiskuserGet(clt *core.SDKClient, req *antifraud.TaobaoAntifraudRiskuserGetAPIRequest, session string) (*antifraud.TaobaoAntifraudRiskuserGetAPIResponse, error) {
	var resp antifraud.TaobaoAntifraudRiskuserGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
