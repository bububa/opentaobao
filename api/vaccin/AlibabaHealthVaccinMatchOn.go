package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinmatchon isv自主上下线疫苗，可以选择上线还是下线
// alibaba.health.vaccin.match.on
//
// isv自主上下线疫苗，可以选择上线还是下线
func Alibabahealthvaccinmatchon(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinmatchonAPIRequest, session string) (*vaccin.AlibabahealthvaccinmatchonAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinmatchonAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
