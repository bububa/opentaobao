package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinNoticeUserBind 支付宝疫苗绑定接种人
// alibaba.health.vaccin.notice.user.bind
//
// 支付宝疫苗绑定接种人
func AlibabaHealthVaccinNoticeUserBind(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeUserBindAPIRequest, resp *vaccin.AlibabaHealthVaccinNoticeUserBindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
