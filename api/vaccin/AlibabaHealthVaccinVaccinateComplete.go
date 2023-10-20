package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinvaccinatecomplete 接种完成反馈接口
// alibaba.health.vaccin.vaccinate.complete
//
// ISV 将用户完成接种的疫苗同步给免疫规划中心
func Alibabahealthvaccinvaccinatecomplete(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinvaccinatecompleteAPIRequest, session string) (*vaccin.AlibabahealthvaccinvaccinatecompleteAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinvaccinatecompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
