package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinnoticeordercancel 福州疫苗取消预约
// alibaba.health.vaccin.notice.order.cancel
//
// 福州疫苗用户取消预约接口
func Alibabahealthvaccinnoticeordercancel(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinnoticeordercancelAPIRequest, session string) (*vaccin.AlibabahealthvaccinnoticeordercancelAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinnoticeordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
