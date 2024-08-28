package vaccin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinNoticeUserBind 支付宝疫苗绑定接种人
// alibaba.health.vaccin.notice.user.bind
//
// 支付宝疫苗绑定接种人
func AlibabaHealthVaccinNoticeUserBind(ctx context.Context, clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeUserBindAPIRequest, resp *vaccin.AlibabaHealthVaccinNoticeUserBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
