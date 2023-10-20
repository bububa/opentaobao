package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// Aliyunaliospayperiodagreementpay 周期扣款支付接口
// aliyun.alios.pay.period.agreement.pay
//
// 周期扣款支付接口，商户服务端通过该接口完成后续期数的支付
func Aliyunaliospayperiodagreementpay(clt *core.SDKClient, req *aliospay.AliyunaliospayperiodagreementpayAPIRequest, session string) (*aliospay.AliyunaliospayperiodagreementpayAPIResponse, error) {
	var resp aliospay.AliyunaliospayperiodagreementpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
