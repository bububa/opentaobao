package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// TaobaoPhoneBankCreditProcess 虚拟话费任务银行信用卡办理进度回传
// taobao.phone.bank.credit.process
//
// 虚拟话费任务银行信用卡办理进度回传
func TaobaoPhoneBankCreditProcess(ctx context.Context, clt *core.SDKClient, req *alicom.TaobaoPhoneBankCreditProcessAPIRequest, resp *alicom.TaobaoPhoneBankCreditProcessAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
