package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayPeriodAgreementUnsign 周期扣款协议解约接口
// aliyun.alios.pay.period.agreement.unsign
//
// 周期扣款协议解约接口
func AliyunAliosPayPeriodAgreementUnsign(clt *core.SDKClient, req *aliospay.AliyunAliosPayPeriodAgreementUnsignAPIRequest, session string) (*aliospay.AliyunAliosPayPeriodAgreementUnsignAPIResponse, error) {
	var resp aliospay.AliyunAliosPayPeriodAgreementUnsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
