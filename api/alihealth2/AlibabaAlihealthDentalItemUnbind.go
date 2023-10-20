package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthdentalitemunbind ISV解绑商品
// alibaba.alihealth.dental.item.unbind
//
// ISV解绑商品
func Alibabaalihealthdentalitemunbind(clt *core.SDKClient, req *alihealth2.AlibabaalihealthdentalitemunbindAPIRequest, session string) (*alihealth2.AlibabaalihealthdentalitemunbindAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthdentalitemunbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
