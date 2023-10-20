package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayPeriodAgreementPay 周期扣款支付接口
// aliyun.alios.pay.period.agreement.pay
//
// 周期扣款支付接口，商户服务端通过该接口完成后续期数的支付
func AliyunAliosPayPeriodAgreementPay(clt *core.SDKClient, req *aliospay.AliyunAliosPayPeriodAgreementPayAPIRequest, resp *aliospay.AliyunAliosPayPeriodAgreementPayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
