package vaccin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinNoticeReplantRemind 支付宝疫苗补种提醒信息
// alibaba.health.vaccin.notice.replant.remind
//
// 支付宝疫苗补种提醒
func AlibabaHealthVaccinNoticeReplantRemind(ctx context.Context, clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeReplantRemindAPIRequest, resp *vaccin.AlibabaHealthVaccinNoticeReplantRemindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
