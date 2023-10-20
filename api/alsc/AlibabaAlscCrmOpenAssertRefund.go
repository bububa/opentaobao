package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmopenassertrefund 资产核销回退接口
// alibaba.alsc.crm.open.assert.refund
//
// 回退已经核销的储值积分券资产
func Alibabaalsccrmopenassertrefund(clt *core.SDKClient, req *alsc.AlibabaalsccrmopenassertrefundAPIRequest, session string) (*alsc.AlibabaalsccrmopenassertrefundAPIResponse, error) {
	var resp alsc.AlibabaalsccrmopenassertrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
