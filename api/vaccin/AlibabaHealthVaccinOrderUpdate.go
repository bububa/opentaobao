package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinOrderUpdate 回传取号状态
// alibaba.health.vaccin.order.update
//
// 回传取号状态
func AlibabaHealthVaccinOrderUpdate(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinOrderUpdateAPIRequest, resp *vaccin.AlibabaHealthVaccinOrderUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
