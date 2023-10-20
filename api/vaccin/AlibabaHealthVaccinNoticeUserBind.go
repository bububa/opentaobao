package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinnoticeuserbind 支付宝疫苗绑定接种人
// alibaba.health.vaccin.notice.user.bind
//
// 支付宝疫苗绑定接种人
func Alibabahealthvaccinnoticeuserbind(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinnoticeuserbindAPIRequest, session string) (*vaccin.AlibabahealthvaccinnoticeuserbindAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinnoticeuserbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
