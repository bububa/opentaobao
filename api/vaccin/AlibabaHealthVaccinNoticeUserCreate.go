package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinnoticeusercreate 支付宝医疗健康疫苗用户创建
// alibaba.health.vaccin.notice.user.create
//
// 支付宝医疗健康疫苗用户创建
func Alibabahealthvaccinnoticeusercreate(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinnoticeusercreateAPIRequest, session string) (*vaccin.AlibabahealthvaccinnoticeusercreateAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinnoticeusercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
