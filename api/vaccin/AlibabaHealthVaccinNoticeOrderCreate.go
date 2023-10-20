package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinnoticeordercreate 支付宝医疗健康疫苗预约创建
// alibaba.health.vaccin.notice.order.create
//
// 支付宝医疗健康疫苗预约创建
func Alibabahealthvaccinnoticeordercreate(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinnoticeordercreateAPIRequest, session string) (*vaccin.AlibabahealthvaccinnoticeordercreateAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinnoticeordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
