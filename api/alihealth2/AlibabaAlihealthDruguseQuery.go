package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthdrugusequery 合理用药规则查询
// alibaba.alihealth.druguse.query
//
// 查询用户购买的药品命中的风险规则
func Alibabaalihealthdrugusequery(clt *core.SDKClient, req *alihealth2.AlibabaalihealthdrugusequeryAPIRequest, session string) (*alihealth2.AlibabaalihealthdrugusequeryAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthdrugusequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
