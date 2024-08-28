package vaccin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinNoticeOrderSign 福州疫苗签到成功通知
// alibaba.health.vaccin.notice.order.sign
//
// 福州疫苗用户签到成功记录
func AlibabaHealthVaccinNoticeOrderSign(ctx context.Context, clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeOrderSignAPIRequest, resp *vaccin.AlibabaHealthVaccinNoticeOrderSignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
