package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeadvisersubmitaccount 顾问入职离职
// alibaba.alihouse.newhome.adviser.submit.account
//
// 顾问入职离职
func Alibabaalihousenewhomeadvisersubmitaccount(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeadvisersubmitaccountAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeadvisersubmitaccountAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeadvisersubmitaccountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
