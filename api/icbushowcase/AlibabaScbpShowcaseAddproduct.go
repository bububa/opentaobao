package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// Alibabascbpshowcaseaddproduct 批量添加橱窗商品
// alibaba.scbp.showcase.addproduct
//
// 批量添加商品到橱窗
func Alibabascbpshowcaseaddproduct(clt *core.SDKClient, req *icbushowcase.AlibabascbpshowcaseaddproductAPIRequest, session string) (*icbushowcase.AlibabascbpshowcaseaddproductAPIResponse, error) {
	var resp icbushowcase.AlibabascbpshowcaseaddproductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
