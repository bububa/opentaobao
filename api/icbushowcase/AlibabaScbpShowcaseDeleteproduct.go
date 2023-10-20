package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// Alibabascbpshowcasedeleteproduct 批量删除橱窗商品
// alibaba.scbp.showcase.deleteproduct
//
// 批量删除橱窗商品
func Alibabascbpshowcasedeleteproduct(clt *core.SDKClient, req *icbushowcase.AlibabascbpshowcasedeleteproductAPIRequest, session string) (*icbushowcase.AlibabascbpshowcasedeleteproductAPIResponse, error) {
	var resp icbushowcase.AlibabascbpshowcasedeleteproductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
