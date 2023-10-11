package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayPeriodAgreementPay 周期扣款支付接口
// aliyun.alios.pay.period.agreement.pay
//
// 周期扣款支付接口，商户服务端通过该接口完成后续期数的支付
func AliyunAliosPayPeriodAgreementPay(clt *core.SDKClient, req *aliospay.AliyunAliosPayPeriodAgreementPayAPIRequest, session string) (*aliospay.AliyunAliosPayPeriodAgreementPayAPIResponse, error) {
	var resp aliospay.AliyunAliosPayPeriodAgreementPayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
