package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayPeriodAgreementStatusGet 查询周期扣款签约状态
// aliyun.alios.pay.period.agreement.status.get
//
// 查询周期扣款签约状态
func AliyunAliosPayPeriodAgreementStatusGet(clt *core.SDKClient, req *aliospay.AliyunAliosPayPeriodAgreementStatusGetAPIRequest, session string) (*aliospay.AliyunAliosPayPeriodAgreementStatusGetAPIResponse, error) {
	var resp aliospay.AliyunAliosPayPeriodAgreementStatusGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
