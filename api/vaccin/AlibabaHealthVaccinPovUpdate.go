package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinPovUpdate 新增/变更接种点信息
// alibaba.health.vaccin.pov.update
//
// ISV 将疫苗的接种点信息同步到免疫规划中心，提醒用户接种时可提供接种点详情。
func AlibabaHealthVaccinPovUpdate(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinPovUpdateAPIRequest, resp *vaccin.AlibabaHealthVaccinPovUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
