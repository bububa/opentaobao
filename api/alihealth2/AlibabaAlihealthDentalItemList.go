package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthdentalitemlist ISV获取口腔标品列表
// alibaba.alihealth.dental.item.list
//
// ISV获取口腔标品列表
func Alibabaalihealthdentalitemlist(clt *core.SDKClient, req *alihealth2.AlibabaalihealthdentalitemlistAPIRequest, session string) (*alihealth2.AlibabaalihealthdentalitemlistAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthdentalitemlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
