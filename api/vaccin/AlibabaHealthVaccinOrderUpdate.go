package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// AlibabaHealthVaccinOrderUpdate 回传取号状态
// alibaba.health.vaccin.order.update
//
// 回传取号状态
func AlibabaHealthVaccinOrderUpdate(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinOrderUpdateAPIRequest, session string) (*vaccin.AlibabaHealthVaccinOrderUpdateAPIResponse, error) {
	var resp vaccin.AlibabaHealthVaccinOrderUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
