package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinnoticereplantremind 支付宝疫苗补种提醒信息
// alibaba.health.vaccin.notice.replant.remind
//
// 支付宝疫苗补种提醒
func Alibabahealthvaccinnoticereplantremind(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinnoticereplantremindAPIRequest, session string) (*vaccin.AlibabahealthvaccinnoticereplantremindAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinnoticereplantremindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
