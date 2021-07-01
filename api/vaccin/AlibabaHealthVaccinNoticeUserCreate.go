package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

/* AlibabaHealthVaccinNoticeUserCreate
支付宝医疗健康疫苗用户创建
alibaba.health.vaccin.notice.user.create

支付宝医疗健康疫苗用户创建 */
func AlibabaHealthVaccinNoticeUserCreate(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinNoticeUserCreateAPIRequest, session string) (*vaccin.AlibabaHealthVaccinNoticeUserCreateAPIResponse, error) {
	var resp vaccin.AlibabaHealthVaccinNoticeUserCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
