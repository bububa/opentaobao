package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoServindustryFinanceInsuranceAgreementSign 保司合同签约后回调接口
// taobao.servindustry.finance.insurance.agreement.sign
//
// 保司合同签约后回调接口
func TaobaoServindustryFinanceInsuranceAgreementSign(ctx context.Context, clt *core.SDKClient, req *trade.TaobaoServindustryFinanceInsuranceAgreementSignAPIRequest, resp *trade.TaobaoServindustryFinanceInsuranceAgreementSignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
