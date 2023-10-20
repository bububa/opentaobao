package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmopenassertverify 资产核销接口
// alibaba.alsc.crm.open.assert.verify
//
// 核销储值，积分，券资产
func Alibabaalsccrmopenassertverify(clt *core.SDKClient, req *alsc.AlibabaalsccrmopenassertverifyAPIRequest, session string) (*alsc.AlibabaalsccrmopenassertverifyAPIResponse, error) {
	var resp alsc.AlibabaalsccrmopenassertverifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
