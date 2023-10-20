package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// TaobaoPhoneBankCreditCheck 虚拟话费任务银行信用卡办理检查
// taobao.phone.bank.credit.check
//
// 虚拟话费任务银行信用卡办理检查
func TaobaoPhoneBankCreditCheck(clt *core.SDKClient, req *alicom.TaobaoPhoneBankCreditCheckAPIRequest, resp *alicom.TaobaoPhoneBankCreditCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
