package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayPeriodAgreementUnsign 周期扣款协议解约接口
// aliyun.alios.pay.period.agreement.unsign
//
// 周期扣款协议解约接口
func AliyunAliosPayPeriodAgreementUnsign(clt *core.SDKClient, req *aliospay.AliyunAliosPayPeriodAgreementUnsignAPIRequest, resp *aliospay.AliyunAliosPayPeriodAgreementUnsignAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
