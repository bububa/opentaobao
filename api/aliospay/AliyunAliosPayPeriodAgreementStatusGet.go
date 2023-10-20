package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayPeriodAgreementStatusGet 查询周期扣款签约状态
// aliyun.alios.pay.period.agreement.status.get
//
// 查询周期扣款签约状态
func AliyunAliosPayPeriodAgreementStatusGet(clt *core.SDKClient, req *aliospay.AliyunAliosPayPeriodAgreementStatusGetAPIRequest, resp *aliospay.AliyunAliosPayPeriodAgreementStatusGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
