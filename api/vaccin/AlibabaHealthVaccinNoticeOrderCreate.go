package vaccin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinNoticeOrderCreate 支付宝医疗健康疫苗预约创建
// alibaba.health.vaccin.notice.order.create
//
// 支付宝医疗健康疫苗预约创建
func AlibabaHealthVaccinNoticeOrderCreate(ctx context.Context, clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeOrderCreateAPIRequest, resp *vaccin.AlibabaHealthVaccinNoticeOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
