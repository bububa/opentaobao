package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinorderupdate 回传取号状态
// alibaba.health.vaccin.order.update
//
// 回传取号状态
func Alibabahealthvaccinorderupdate(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinorderupdateAPIRequest, session string) (*vaccin.AlibabahealthvaccinorderupdateAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinorderupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
