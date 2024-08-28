package aliospay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayPeriodAgreementUnsign 周期扣款协议解约接口
// aliyun.alios.pay.period.agreement.unsign
//
// 周期扣款协议解约接口
func AliyunAliosPayPeriodAgreementUnsign(ctx context.Context, clt *core.SDKClient, req *aliospay.AliyunAliosPayPeriodAgreementUnsignAPIRequest, resp *aliospay.AliyunAliosPayPeriodAgreementUnsignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
