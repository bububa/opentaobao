package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// Alibabascbpshowcaseupdateproduct 替换橱窗商品
// alibaba.scbp.showcase.updateproduct
//
// 替换橱窗商品
func Alibabascbpshowcaseupdateproduct(clt *core.SDKClient, req *icbushowcase.AlibabascbpshowcaseupdateproductAPIRequest, session string) (*icbushowcase.AlibabascbpshowcaseupdateproductAPIResponse, error) {
	var resp icbushowcase.AlibabascbpshowcaseupdateproductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
