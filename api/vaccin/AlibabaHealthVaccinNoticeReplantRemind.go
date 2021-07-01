package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

/* AlibabaHealthVaccinNoticeReplantRemind
支付宝疫苗补种提醒信息
alibaba.health.vaccin.notice.replant.remind

支付宝疫苗补种提醒 */
func AlibabaHealthVaccinNoticeReplantRemind(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeReplantRemindAPIRequest, session string) (*vaccin.AlibabaHealthVaccinNoticeReplantRemindAPIResponse, error) {
	var resp vaccin.AlibabaHealthVaccinNoticeReplantRemindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
