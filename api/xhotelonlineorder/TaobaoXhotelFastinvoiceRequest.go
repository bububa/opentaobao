package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelFastinvoiceRequest 极速开票开票请求回传
// taobao.xhotel.fastinvoice.request
//
// 极速开票开票请求回传,用于记录航信开票请求数据
func TaobaoXhotelFastinvoiceRequest(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelFastinvoiceRequestAPIRequest, session string) (*xhotelonlineorder.TaobaoXhotelFastinvoiceRequestAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoXhotelFastinvoiceRequestAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
