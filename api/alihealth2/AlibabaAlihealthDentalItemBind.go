package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthdentalitembind ISV绑定外部门店id和外部商品id
// alibaba.alihealth.dental.item.bind
//
// ISV绑定外部门店id和外部商品id
func Alibabaalihealthdentalitembind(clt *core.SDKClient, req *alihealth2.AlibabaalihealthdentalitembindAPIRequest, session string) (*alihealth2.AlibabaalihealthdentalitembindAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthdentalitembindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
