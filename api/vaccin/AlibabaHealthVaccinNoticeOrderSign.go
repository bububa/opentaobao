package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinNoticeOrderSign 福州疫苗签到成功通知
// alibaba.health.vaccin.notice.order.sign
//
// 福州疫苗用户签到成功记录
func AlibabaHealthVaccinNoticeOrderSign(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeOrderSignAPIRequest, resp *vaccin.AlibabaHealthVaccinNoticeOrderSignAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
