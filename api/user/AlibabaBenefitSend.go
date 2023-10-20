package user

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// Alibababenefitsend 发奖接口
// alibaba.benefit.send
//
// 发奖接口
func Alibababenefitsend(clt *core.SDKClient, req *user.AlibababenefitsendAPIRequest, session string) (*user.AlibababenefitsendAPIResponse, error) {
	var resp user.AlibababenefitsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
